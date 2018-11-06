package main

import (
	"context"
	"flag"
	"fmt"

	ds "github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	logging "github.com/ipfs/go-log"
	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
	common "github.com/prestonvanloon/libp2p-chat/common"
)

func init() {
	logging.SetDebugLogging()
}

var (
	privateKey = flag.String("private", "", "Private key to use for peer ID")
)

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listen, err := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/4001")
	if err != nil {
		panic(err)
	}
	opts := []libp2p.Option{
		libp2p.ListenAddrs(listen),
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

	host, err := libp2p.New(ctx, opts...)
	if err != nil {
		panic(err)
	}

	// Construct a datastore (needed by the DHT). This is just a simple, in-memory thread-safe datastore.
	dstore := dsync.MutexWrap(ds.NewMapDatastore())

	// Make the DHT
	dht := kaddht.NewDHT(ctx, host, dstore)
	if err := dht.Bootstrap(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("DHT node running: /ip4/0.0.0.0/tcp/%v/p2p/%s\n", 4001, host.ID().Pretty())

	// TODO: Log peer connections?

	common.ExportMetrics("github.com/prestonvanloon/libp2p-chat/dht")
}
