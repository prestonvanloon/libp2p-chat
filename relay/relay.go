package main

import (
	"context"
	"flag"
	"fmt"

	logger "github.com/ipfs/go-log"
	libp2p "github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	crypto "github.com/libp2p/go-libp2p-crypto"
	multiaddr "github.com/multiformats/go-multiaddr"
	common "github.com/prestonvanloon/libp2p-chat/common"
)

var (
	privateKey = flag.String("private", "", "Private key to use for peer ID")
	debug      = flag.Bool("v", false, "Verbose logging")
)

func main() {
	flag.Parse()

	if *debug {
		logger.SetDebugLogging()
		fmt.Println("Verbose logging on")
	}

	port := 6660
	srcMAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	opts := []libp2p.Option{
		libp2p.EnableRelay(circuit.OptHop),
		libp2p.ListenAddrs(srcMAddr),
	}

	if *privateKey != "" {
		b, err := crypto.ConfigDecodeKey(*privateKey)
		if err != nil {
			panic(err)
		}
		pk, err := crypto.UnmarshalPrivateKey(b)
		if err != nil {
			panic(err)
		}
		opts = append(opts, libp2p.Identity(pk))
	}

	h, err := libp2p.New(
		context.Background(),
		opts...,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Relay available: /ip4/0.0.0.0/tcp/%v/p2p/%s\n", port, h.ID().Pretty())

	common.ExportMetrics("github.com/prestonvanloon/libp2p-chat/relay")
}
