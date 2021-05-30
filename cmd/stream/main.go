package main

import (
	"flag"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/present-intro-nats/strm"
)

var deleteMode bool

func init() {
	flag.BoolVar(&deleteMode, "delete", false,
		`Deletes example JetStream instead of creating it.
If this flag is not provided, it creates an example stream.`)
	flag.Parse()
}

func main() {
	var nc *nats.Conn
	js := strm.JetStreamContext(nc)
	defer nc.Close()

	const strName = "tst"
	if !deleteMode {
		strm.Create(js, strName)
		return
	}

	strm.Delete(js, strName)
}
