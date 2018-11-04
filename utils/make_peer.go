package utils

import (
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
)

// MakePeer from address string.
func MakePeer(addr string) (*peerstore.PeerInfo, error) {
	maddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	return peerstore.InfoFromP2pAddr(maddr)
}
