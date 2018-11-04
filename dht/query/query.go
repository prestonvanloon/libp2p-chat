package main

import (
	"context"
	"fmt"
	"os"
	"time"

	ggio "github.com/gogo/protobuf/io"
	"github.com/libp2p/go-libp2p"
	dhtopts "github.com/libp2p/go-libp2p-kad-dht/opts"
	dhtpb "github.com/libp2p/go-libp2p-kad-dht/pb"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/prestonvanloon/libp2p-chat/utils"
)

func main() {
	addr := os.Args[1]
	fmt.Println("Querying " + addr)
	queryNode(addr)
	fmt.Printf("Press ctrl+C to stop")
	select {}
}

func queryNode(addr string) {
	h, err := libp2p.New(context.Background())
	if err != nil {
		panic(err)
	}

	info, err := utils.MakePeer(addr)
	if err != nil {
		panic(err)
	}

	h.Connect(context.Background(), *info)

	s, err := h.NewStream(context.Background(), info.ID, dhtopts.ProtocolDHT)
	if err != nil {
		panic(err)
	}

	r := ggio.NewDelimitedReader(s, 2000000)
	w := ggio.NewDelimitedWriter(s)

	go readIncomingMessages(r)

	for {
		sendQuery(w, info.ID)

		time.Sleep(5 * time.Second)
	}
}

func readIncomingMessages(r ggio.ReadCloser) {
	for {
		var resp dhtpb.Message
		if err := r.ReadMsg(&resp); err != nil {
			panic(err)
		}
		fmt.Printf("Msg received: %+v\n", resp)
	}
}

func sendQuery(w ggio.WriteCloser, pid peer.ID) {
	m := dhtpb.NewMessage(dhtpb.Message_FIND_NODE, []byte(pid), 0)
	if err := w.WriteMsg(m); err != nil {
		panic(err)
	}
}
