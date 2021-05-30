package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/present-intro-nats/strm"
)

func main() {
	var nc *nats.Conn
	js := strm.JetStreamContext(nc)
	defer nc.Close()

	cTest := pullSub(js, "test", "test-go") // sets up consumer tst test-go if not exists
	cTestDot := pullSub(js, "test.>", "test-dot")

	showMsgs(cTest)
	showMsgs(cTestDot)

	time.Sleep(3 * time.Second) // allow showMsgs to execute
}

func pullSub(js nats.JetStreamContext, subj, consName string) *nats.Subscription {
	sub, err := js.PullSubscribe(subj, consName)
	if err != nil {
		log.Panicf("could not pull subscribe on subject: %s, consumer name: %s, :%v",
			subj, consName, err)
	}
	return sub
}

func showMsgs(sub *nats.Subscription) {
	go func() {
		for {
			fetch(sub)
		}
	}()
}
func fetch(sub *nats.Subscription) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	const batchSize = 10
	msgs, _ := sub.Fetch(batchSize, nats.Context(ctx))
	for _, msg := range msgs {
		msg.Ack()

		inf, err := sub.ConsumerInfo()
		if err != nil {
			log.Printf("could get get consumer info: %v", err)
		}
		fmt.Printf("%s %s: %s\n", inf.Stream, inf.Name, msg.Data)
	}
}
