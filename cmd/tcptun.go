package cmd 

import (
	"log"
	"net"
	"tcptun/util"
	"time"
)

// Server is a TCP server that takes an incoming request and sends it to another
type Server struct {
	// TCP address to listen on
	Addr string

	// TCP address of target server
	Target string

	// RequestCipher is an optional function that changes the request from a client to the target server.
	RequestCipher func(b *[]byte, key []byte)

	// ResponseCipher is an optional function that changes the response from the target server.
	ResponseCipher func(b *[]byte, key []byte)

	// Encrypt Key
	Key []byte

	// Server mode
	ServerMode bool
}

func (s *Server) ListenAndServe() {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error")
			log.Println(err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	// connects to target server
	rconn, err := net.DialTimeout("tcp", s.Target, 30*time.Second)
	if err != nil {
		log.Println(err)
		return
	}

	if s.ServerMode {
		s.RequestCipher = func(b *[]byte, key []byte) {
			util.Decrypt(b, key)
		}
		s.ResponseCipher = func(b *[]byte, key []byte) {
			util.Encrypt(b, key)
		}
	} else {
		s.RequestCipher = func(b *[]byte, key []byte) {
			util.Encrypt(b, key)
		}
		s.ResponseCipher = func(b *[]byte, key []byte) {
			util.Decrypt(b, key)
		}
	}
	go s.copy(conn, rconn, s.RequestCipher)
	go s.copy(rconn, conn, s.ResponseCipher)
}

// write to dst what it reads from src
func (s *Server) copy(src, dst net.Conn, cipher func(b *[]byte, key []byte)) {
	defer dst.Close()
	defer src.Close()

	buff := make([]byte, 32768)
	for {
		n, err := src.Read(buff)
		if n == 0 || err != nil {
			return
		}

		b := buff[:n]
		if cipher != nil {
			cipher(&b, s.Key)
		}

		_, err = dst.Write(b)
		if err != nil {
			return
		}
	}
}
