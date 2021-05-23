package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/jsm.go"
	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

var nc *nats.Conn

func main() {
	fmt.Println("Create a Stream")
	mgr := jsmMgr()
	defer nc.Close()

	const maxDur = 1<<63 - 1
	st, err := mgr.NewStream("tst", jsm.Subjects("test.>"),
		jsm.MaxAge(maxDur), jsm.FileStorage())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}

	fmt.Println("Stream Info")
	inf, err := st.Information()
	if err != nil {
		log.Fatalf("could not get stream info: %v", err)
	}
	prettyPrint(inf)

	fmt.Println("Delete a Stream")
	st.Delete()
}

func jsmMgr() *jsm.Manager {
	var err error
	natsURL := dflt.EnvString("NATS_URL", "nats://127.0.0.1:4222")
	nc, err = nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("could not connect to NATS: %v", err)
	}

	mgr, err := jsm.New(nc)
	if err != nil {
		log.Fatal("could not create jetstream manager: %v", err)
	}

	return mgr
}

func prettyPrint(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatalf("could not prettyPrint: %v", err)
	}
	fmt.Println(string(b))
}
