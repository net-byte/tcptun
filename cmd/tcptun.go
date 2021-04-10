package cmd

import (
	"log"
	"net"
	"time"

	"github.com/net-byte/tcptun/common/cipher"
)

// Server is a secure TCP proxy server
type Server struct {
	// TCP address of local server
	LocalAddr string

	// TCP address of target server
	ServerAddr string

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
	ln, err := net.Listen("tcp", s.LocalAddr)
	if err != nil {
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

	if s.ServerMode {
		s.RequestCipher = func(b *[]byte, key []byte) {
			cipher.Decrypt(b, key)
		}
		s.ResponseCipher = func(b *[]byte, key []byte) {
			cipher.Encrypt(b, key)
		}
	} else {
		s.RequestCipher = func(b *[]byte, key []byte) {
			cipher.Encrypt(b, key)
		}
		s.ResponseCipher = func(b *[]byte, key []byte) {
			cipher.Decrypt(b, key)
		}
	}
	go s.copy(conn, remoteConn, s.RequestCipher)
	go s.copy(remoteConn, conn, s.ResponseCipher)
}

func (s *Server) copy(src, dst net.Conn, cipher func(b *[]byte, key []byte)) {
	defer dst.Close()
	defer src.Close()

	buff := make([]byte, 32768)
	for {
		n, err := src.Read(buff)
		if n == 0 || err != nil {
			break
		}

		b := buff[:n]
		if cipher != nil {
			cipher(&b, s.Key)
		}

		_, err = dst.Write(b)
		if err != nil {
			break
		}
	}
}
