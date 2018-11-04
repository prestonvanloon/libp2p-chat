package main

import (
	"crypto/rand"
	"fmt"

	"github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

func main() {
	// Generate a key, print out the b64 representation of private key and public
	// key.

	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"***************************************************\n"+
			"Peer ID: %v\n"+
			"\n"+
			"Public: %v\n"+
			"\n"+
			"Private: %v\n"+
			"***************************************************\n",
		IDString(prvKey),
		keyString(prvKey.GetPublic()),
		keyString(prvKey),
	)
}

type key interface {
	Bytes() ([]byte, error)
}

func keyString(k key) string {
	b, err := k.Bytes()
	if err != nil {
		panic(err)
	}
	return crypto.ConfigEncodeKey(b)
}

func IDString(k crypto.PrivKey) string {
	pid, err := peer.IDFromPrivateKey(k)
	if err != nil {
		panic(err)
	}
	return pid.Pretty()
}
