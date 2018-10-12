package main

import (
	"context"
	"crypto/rand"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	crypto "github.com/libp2p/go-libp2p-crypto"
	multiaddr "github.com/multiformats/go-multiaddr"
)

func main() {
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		panic(err)
	}

	port := 6660
	srcMAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	h, err := libp2p.New(
		context.Background(),
		libp2p.EnableRelay(circuit.OptHop),
		libp2p.Identity(prvKey),
		libp2p.ListenAddrs(srcMAddr),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Relay available: /ip4/127.0.0.1/tcp/%v/p2p/%s\n", port, h.ID().Pretty())

	// Hang forever.
	select {}
}
