package main

import (
	"github.com/nats-io/nats.go"
	"github.com/siuyin/present-intro-nats/strm"
)

func main() {
	var nc *nats.Conn
	js := strm.JetStreamContext(nc)
	defer nc.Close()

	const strName = "tst"
	//strm.AddConsumer(js, strName, "test-dot", "test.>")
	strm.AddConsumer(js, strName, "test", "test")
}
