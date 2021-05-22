package cmd

import (
	"crypto/rc4"
	"io"
	"log"
	"net"
	"time"
)

// Server is a secure TCP proxy server
type Server struct {
	// TCP address of local server
	LocalAddr string

	// TCP address of target server
	ServerAddr string

	// Encryption Key
	Key []byte
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.LocalAddr)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	remoteConn, err := net.DialTimeout("tcp", s.ServerAddr, 30*time.Second)
	if err != nil {
		log.Println(err)
		return
	}

	go s.copy(conn, remoteConn)
	go s.copy(remoteConn, conn)
}

func (s *Server) copy(src, dst net.Conn) {
	defer dst.Close()
	defer src.Close()
	c, err := rc4.NewCipher(s.Key)
	if err != nil {
		log.Fatalln(err)
	}
	buff := make([]byte, 4096)
	for {
		n, err := src.Read(buff)
		if err != nil || err == io.EOF {
			break
		}
		b := buff[:n]
		c.XORKeyStream(b, b)
		_, err = dst.Write(b)
		if err != nil {
			break
		}
	}
}
