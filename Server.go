package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Server struct {
	clients map[string]net.Conn
	mu      sync.Mutex
}

func NewServer() *Server {
	return &Server{clients: make(map[string]net.Conn)}
}
func (s *Server) Run() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Server is running on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	s.mu.Lock()
	s.clients[username] = conn
	s.mu.Unlock()
	fmt.Printf("%s has joined the chat\n", username)
	s.broadcast(&Message{From: "Server", Content: username + " joined"}, username)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		msg, err := (&Message{}).ToString(line)
		if err != nil {
			continue
		}
		if msg.To == "" {
			s.broadcast(msg, username)
		} else {
			s.sendTo(msg, msg.To)
		}
	}
	s.mu.Lock()
	delete(s.clients, username)
	s.mu.Unlock()
	fmt.Println(username, "left the chat")
	s.broadcast(&Message{From: "Server", Content: username + " left"}, username)
	defer conn.Close()
}

func (s *Server) broadcast(msg *Message, except string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for user, conn := range s.clients {
		if user != except {
			conn.Write(msg.ToJSON())
		}
	}
}

func (s *Server) sendTo(msg *Message, to string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if conn, ok := s.clients[to]; ok {
		conn.Write(msg.ToJSON())
	}
}
