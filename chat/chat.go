/*
 *
 * The MIT License (MIT)
 *
 * Copyright (c) 2014 Juan Batiz-Benet
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * This program demonstrate a simple chat application using p2p communication.
 *
 */

package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"os"
	"time"

	"github.com/libp2p/go-libp2p"

	ds "github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	logger "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	"github.com/prestonvanloon/libp2p-chat/utils"
)

var outgoingStreams []peer.ID
var tellTheTime = flag.Bool("tell_time", false, "For each stream, tell the time every 5 seconds")

func handleStream(s net.Stream) {

	log.Printf("Got a new stream from %s via %s", s.Conn().RemotePeer().Pretty(), s.Conn().RemoteMultiaddr().String())

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw)
	go writeData(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}
func readData(rw *bufio.ReadWriter) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
		}

	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	if *tellTheTime {
		go func() {
			for {
				rw.WriteString(fmt.Sprintf("The current time in k8s is %s", time.Now()))
				time.Sleep(5 * time.Second)
			}
		}()
	}

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		rw.WriteString(fmt.Sprintf("%s\n", sendData))
		rw.Flush()
	}

}

func main() {
	sourcePort := flag.Int("sp", 0, "Source port number")
	dest := flag.String("d", "", "Destination multiaddr string")
	help := flag.Bool("help", false, "Display help")
	debug := flag.Bool("debug", false, "Debug generates the same node ID on every execution")
	relay := flag.String("relay", "", "Relay node to connect")
	dhtAddr := flag.String("dht", "", "DHT node to retrieve peers")
	v := flag.Bool("v", false, "Verbose logging (debug level)")
	flag.Parse()

	if *v {
		logger.SetDebugLogging()
	}

	if *help {
		fmt.Printf("This program demonstrates a simple p2p chat application using libp2p\n\n")
		fmt.Println("Usage: Run './chat -sp <SOURCE_PORT>' where <SOURCE_PORT> can be any port number.")
		fmt.Println("Now run './chat -d <MULTIADDR>' where <MULTIADDR> is multiaddress of previous listener host.")

		os.Exit(0)
	}

	// If debug is enabled, use a constant random source to generate the peer ID. Only useful for debugging,
	// off by default. Otherwise, it uses rand.Reader.
	var r io.Reader
	if *debug {
		// Use the port number as the randomness source.
		// This will always generate the same host ID on multiple executions, if the same port number is used.
		// Never do this in production code.
		r = mrand.New(mrand.NewSource(int64(*sourcePort)))
	} else {
		r = rand.Reader
	}

	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		panic(err)
	}

	// 0.0.0.0 will listen on any interface device.
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", *sourcePort))

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	h, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
		libp2p.EnableRelay(),
		libp2p.AddrsFactory(addRelayAddrs(*relay, true /*relayOnly*/)),
	)

	if err != nil {
		panic(err)
	}

	// Connect to relay
	if *relay != "" {
		info, err := utils.MakePeer(*relay)
		if err != nil {
			panic(err)
		}

		if err := h.Connect(context.Background(), *info); err != nil {
			panic(err)
		}
		fmt.Println("Connected to relay " + info.ID.Pretty())
	}

	// Connect to DHT
	if *dhtAddr != "" {
		dht := kaddht.NewDHT(context.Background(), h, dsync.MutexWrap(ds.NewMapDatastore()))
		h = rhost.Wrap(h, dht)

		info, err := utils.MakePeer(*dhtAddr)
		if err != nil {
			panic(err)
		}

		if err := h.Connect(context.Background(), *info); err != nil {
			panic(err)
		}
		fmt.Println("Connected to DHT " + info.ID.Pretty())
		dht.Bootstrap(context.Background())
	}

	if *dest == "" {
		// Set a function as stream handler.
		// This function is called when a peer connects, and starts a stream with this protocol.
		// Only applies on the receiving side.
		h.SetStreamHandler("/chat/1.0.0", handleStream)

		// Let's get the actual TCP port from our listen multiaddr, in case we're using 0 (default; random available port).
		var port string
		for _, la := range h.Network().ListenAddresses() {
			if p, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
				port = p
				break
			}
		}

		if port == "" {
			panic("was not able to find actual local port")
		}

		go func() {
			for {
				for _, pid := range h.Peerstore().Peers() {
					//fmt.Printf("Connected peer: %v\n", pid.Pretty())
					protocols, err := h.Peerstore().GetProtocols(pid)
					if err != nil {
						panic(err)
					}
					if containsProtocol(protocols, "/chat/1.0.0") && !containsPeer(outgoingStreams, pid) {
						// Connect to this peer!!
						startStreamWithPeer(h, pid)
					}
				}
				//fmt.Println("------------------------")
				time.Sleep(5 * time.Second)
			}
		}()

		fmt.Printf("Run './chat -d /ip4/127.0.0.1/tcp/%v/p2p/%s' on another console.\n", port, h.ID().Pretty())
		fmt.Println("You can replace 127.0.0.1 with public IP as well.")
		fmt.Printf("\nWaiting for incoming connection\n\n")

		// Hang forever
		<-make(chan struct{})
	} else {
		fmt.Println("This node's multiaddresses:")
		for _, la := range h.Addrs() {
			fmt.Printf(" - %v\n", la)
		}
		fmt.Println()

		// Turn the destination into a multiaddr.
		maddr, err := multiaddr.NewMultiaddr(*dest)
		if err != nil {
			log.Fatalln(err)
		}

		// Extract the peer ID from the multiaddr.
		info, err := peerstore.InfoFromP2pAddr(maddr)
		if err != nil {
			log.Fatalln(err)
		}

		if *relay != "" {
			relayaddr, err := multiaddr.NewMultiaddr("/p2p-circuit/ipfs/" + info.ID.Pretty())
			if err != nil {
				panic(err)
			}
			info.Addrs = []multiaddr.Multiaddr{relayaddr}
		}

		// Add the destination's peer multiaddress in the peerstore.
		// This will be used during connection and stream creation by libp2p.
		h.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

		startStreamWithPeer(h, info.ID)

		// Hang forever.
		select {}
	}
}

func startStreamWithPeer(h host.Host, pid peer.ID) {
	outgoingStreams = append(outgoingStreams, pid)
	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
	s, err := h.NewStream(context.Background(), pid, "/chat/1.0.0")
	if err != nil {
		panic(err)
	}

	// Create a buffered stream so that read and writes are non blocking.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// Create a thread to read and write data.
	go writeData(rw)
	go readData(rw)
}
