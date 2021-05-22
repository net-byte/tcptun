package main

import (
	"flag"
	"log"

	"github.com/net-byte/tcptun/cmd"
)

var (
	localAddr  = flag.String("l", ":2000", "local address")
	serverAddr = flag.String("s", ":2001", "server address")
	key        = flag.String("k", "NcRfWjXn3r4u7x", "encryption key")
)

func main() {
	flag.Parse()

	s := cmd.Server{
		LocalAddr:  *localAddr,
		ServerAddr: *serverAddr,
		Key:        []byte(*key),
	}

	log.Println("Proxying from " + s.LocalAddr + " to " + s.ServerAddr)
	s.Start()
}
