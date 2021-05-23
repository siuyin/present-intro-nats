package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/jsm.go"
	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

var nc *nats.Conn

func main() {
	fmt.Println("Exercise a consumer")
	mgr := jsmMgr()
	defer nc.Close()

	st := tstStream(mgr)
	defer st.Delete()
	pubDat("test.x")

	cons := xConsumer(mgr, st.Name())
	consumeDat(st, cons)

}

func jsmMgr() *jsm.Manager {
	var err error
	natsURL := dflt.EnvString("NATS_URL", "nats://127.0.0.1:4222")
	nc, err = nats.Connect(natsURL, nats.UseOldRequestStyle())
	if err != nil {
		log.Fatalf("could not connect to NATS: %v", err)
	}

	mgr, err := jsm.New(nc)
	if err != nil {
		log.Fatal("could not create jetstream manager: %v", err)
	}

	return mgr
}

func tstStream(mgr *jsm.Manager) *jsm.Stream {
	const maxDur = 1<<63 - 1
	st, err := mgr.NewStream("tst", jsm.Subjects("test.>"),
		jsm.MaxAge(maxDur), jsm.FileStorage())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}

	return st
}

func pubDat(subj string) {
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Msg %d", i)
		nc.Publish(subj, []byte(msg))
	}
}

func xConsumer(mgr *jsm.Manager, stream string) *jsm.Consumer {
	c, err := mgr.NewConsumer(stream,
		jsm.DurableName("x"),
	)
	if err != nil {
		log.Fatalf("could not create consumer: %v", err)
	}
	return c
}

func consumeDat(st *jsm.Stream, cons *jsm.Consumer) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		m, err := cons.NextMsg()
		if err != nil {
			log.Printf("could not get next message: %v", err)
			return
		}
		fmt.Printf("%s\n", m.Data)
	}
	dur := time.Now().Sub(start)
	fmt.Printf("Message processing time: %v ms\n", dur.Milliseconds())
}
func strInfo(st *jsm.Stream) {
	fmt.Println("Stream Info")
	inf, err := st.Information()
	if err != nil {
		log.Fatalf("could not get stream info: %v", err)
	}
	prettyPrint(inf)
}
func prettyPrint(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatalf("could not prettyPrint: %v", err)
	}
	fmt.Println(string(b))
}
