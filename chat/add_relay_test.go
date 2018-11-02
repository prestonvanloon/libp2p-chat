package main

import (
	"testing"

	ma "github.com/multiformats/go-multiaddr"
)

func TestAddRelayAddrs(t *testing.T) {
	relay := "/ip4/127.0.0.1/tcp/6660/p2p/QmQ7zhY7nGY66yK1n8hLGevfVyjbtvHSgtZuXkCH9oTrgi"
	f := addRelayAddrs(relay, false /*relayOnly*/)

	a, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/33201/p2p/QmaXZhW44pwQxBSeLkE5FNeLz8tGTTEsRciFg1DNWXXrWG")
	if err != nil {
		t.Fatal(err)
	}
	addrs := []ma.Multiaddr{a}

	result := f(addrs)

	if len(result) != 2 {
		t.Errorf("Unexpected number of addresses. Wanted %d, got %d", 2, len(result))
	}

	if result[1].String() != "/ip4/127.0.0.1/tcp/6660/ipfs/QmQ7zhY7nGY66yK1n8hLGevfVyjbtvHSgtZuXkCH9oTrgi/p2p-circuit/ipfs/QmaXZhW44pwQxBSeLkE5FNeLz8tGTTEsRciFg1DNWXXrWG" {
		t.Errorf("Address at index 1 (%s) is not the expected p2p-circuit address", result[1].String())
	}
}
