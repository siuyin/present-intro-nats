package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

func main() {
	fmt.Println("Hello JetStream")

	natsURL := dflt.EnvString("NATS_URL", "localhost:4222")
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JetStream context:", js)
}
