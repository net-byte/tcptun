package main

import (
	"flag"
	"log"

	"github.com/net-byte/tcptun/cmd"
	"github.com/net-byte/tcptun/common/cipher"
)

var (
	localAddr  = flag.String("l", ":2000", "Local address")
	serverAddr = flag.String("s", ":2001", "Server address")
	serverMode = flag.Bool("S", false, "Server mode")
	key        = flag.String("k", "6da62287-979a-4eb4-a5ab-8b3d89da134b", "Encrypt key")
)

func main() {
	flag.Parse()

	s := cmd.Server{
		LocalAddr:  *localAddr,
		ServerAddr: *serverAddr,
		ServerMode: *serverMode,
		Key:        cipher.CreateHash(*key),
	}

	log.Println("Proxying from " + s.LocalAddr + " to " + s.ServerAddr)
	s.Start()
}
