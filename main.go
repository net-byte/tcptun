package main

import (
	"flag"
	"log"
	"tcptun/util"

	"tcptun/cmd"
)

var (
	localAddr  = flag.String("l", ":1987", "Proxy local address")
	serverAddr = flag.String("s", ":1080", "Proxy server address")
	serverMode = flag.Bool("S", false, "Server mode")
	key        = flag.String("k", "6da62287-979a-4eb4-a5ab-8b3d89da134b", "Encrypt key")
)

func main() {
	flag.Parse()

	p := cmd.Server{
		Addr:       *localAddr,
		Target:     *serverAddr,
		Key:        util.CreateHash(*key),
		ServerMode: *serverMode,
	}

	log.Println("Proxying from " + p.Addr + " to " + p.Target)
	p.Start()
}
