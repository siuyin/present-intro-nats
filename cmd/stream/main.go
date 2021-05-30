package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

var (
	nc         *nats.Conn
	deleteMode bool
)

func init() {
	flag.BoolVar(&deleteMode, "delete", false,
		`Deletes example JetStream instead of creating it.
If this flag is not provided, it creates an example stream.`)
	flag.Parse()
}

func main() {
	js := jsCtx()
	defer nc.Close()

	const strName = "tst"
	if !deleteMode {
		createStream(js, strName)
		return
	}
	deleteStream(js, strName)
}

func jsCtx() nats.JetStreamContext {
	var err error
	natsURL := dflt.EnvString("NATS_URL", "nats://127.0.0.1:4222")
	nc, err = nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("could not connect to NATS: %v", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("could not create JetStream context: %v", err)
	}

	return js
}

func createStream(js nats.JetStreamContext, name string) *nats.StreamInfo {
	fmt.Printf("Creating stream: %q\n", name)
	strInfo, err := js.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{"test.>", "test"},
		MaxAge:   0, // 0 means keep forever
		Storage:  nats.FileStorage,
	})
	if err != nil {
		log.Panicf("could not create stream: %v", err)
	}

	prettyPrint(strInfo)
	return strInfo
}

func prettyPrint(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatalf("could not prettyPrint: %v", err)
	}
	fmt.Println(string(b))
}

func deleteStream(js nats.JetStreamContext, name string) {
	fmt.Printf("Deleting stream: %q\n", name)
	if err := js.DeleteStream(name); err != nil {
		log.Printf("error deleting stream: %v", err)
	}
}
