package cmd

import (
	"log"
	"net"
	"tcptun/util"
	"time"
)

// Server is a secure TCP proxy server
type Server struct {
	// TCP address of local server
	Addr string

	// TCP address of target server
	Target string

	// RequestCipher
	RequestCipher func(b *[]byte, key []byte)

	// ResponseCipher
	ResponseCipher func(b *[]byte, key []byte)

	// Encrypt Key
	Key []byte

	// Server mode
	ServerMode bool
}

func (s *Server) Start() {
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
