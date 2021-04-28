package cmd

import (
	"io"
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
	RequestCipher func(b []byte) []byte

	// ResponseCipher
	ResponseCipher func(b []byte) []byte

	// Encryption Key
	Key string

	// Server mode
	ServerMode bool
}

func (s *Server) Start() {
	cipher.GenerateKey(s.Key)
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

	if s.ServerMode {
		s.RequestCipher = func(b []byte) []byte {
			return cipher.Decrypt(b)
		}
		s.ResponseCipher = func(b []byte) []byte {
			return cipher.Encrypt(b)
		}
	} else {
		s.RequestCipher = func(b []byte) []byte {
			return cipher.Encrypt(b)
		}
		s.ResponseCipher = func(b []byte) []byte {
			return cipher.Decrypt(b)
		}
	}
	go s.copy(conn, remoteConn, s.RequestCipher)
	go s.copy(remoteConn, conn, s.ResponseCipher)
}

func (s *Server) copy(src, dst net.Conn, cipher func(b []byte) []byte) {
	defer dst.Close()
	defer src.Close()

	buff := make([]byte, 4096)
	for {
		n, err := src.Read(buff)
		if err != nil || err == io.EOF {
			break
		}

		b := buff[:n]
		if cipher != nil {
			b = cipher(b)
		}

		_, err = dst.Write(b)
		if err != nil {
			break
		}
	}
}
