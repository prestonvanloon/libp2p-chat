package main

import (
	"context"
	"flag"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	crypto "github.com/libp2p/go-libp2p-crypto"
	multiaddr "github.com/multiformats/go-multiaddr"
)

var (
	privateKey = flag.String("private", "", "Private key to use for peer ID")
)

func main() {
	flag.Parse()

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

	fmt.Printf("Relay available: /ip4/127.0.0.1/tcp/%v/p2p/%s\n", port, h.ID().Pretty())

	// Hang forever.
	select {}
}
