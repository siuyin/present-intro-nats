package main

import (
	"fmt"
	"log"

	"github.com/nats-io/jsm.go"
	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

func main() {
	fmt.Println("Hello Jetstream")

	natsURL := dflt.EnvString("NATS_URL", "localhost:4222")
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	mgr, err := jsm.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mgr)
}
