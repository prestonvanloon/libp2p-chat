package main

import (
	peer "github.com/libp2p/go-libp2p-peer"
)

func containsProtocol(protocols []string, protocol string) bool {
	for _, p := range protocols {
		if p == protocol {
			return true
		}
	}
	return false
}

func containsPeer(peers []peer.ID, pid peer.ID) bool {
	for _, p := range peers {
		if p == pid {
			return true
		}
	}
	return false
}
