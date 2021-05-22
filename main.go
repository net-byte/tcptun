package main

import (
	"flag"
	"log"

	"github.com/net-byte/tcptun/cmd"
)

var (
	localAddr  = flag.String("l", ":2000", "local address")
	serverAddr = flag.String("s", ":2001", "server address")
	serverMode = flag.Bool("S", false, "server mode")
	key        = flag.String("k", "3tG*Cy%Zt6GWZV8W", "encryption key")
)

func main() {
	flag.Parse()

	s := cmd.Server{
		LocalAddr:  *localAddr,
		ServerAddr: *serverAddr,
		ServerMode: *serverMode,
		Key:        []byte(*key),
	}

	log.Println("Proxying from " + s.LocalAddr + " to " + s.ServerAddr)
	s.Start()
}
