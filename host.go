package main

import (
	"encoding/json"
	"log"
	"net"
	"sync"
)

const (
	skip = iota
	connect
	leave
)

type Req struct {
	Nick string `json:"nick"`
	Opt  int    `json:"opt"`
}

func hostFunc() {
	var clients []string
	var m sync.Mutex
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("Unable to listen on 42069: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Unable to accept connection: %v", err)
		}
		go handleConn(conn, clients, &m)
	}
}

func handleConn(conn net.Conn, clients []string, m *sync.Mutex) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Printf("Unable to read req: %v", err)
	}
	var req Req
	err = json.Unmarshal(buf, &req)
	if err != nil {
		log.Fatalf("Failed to unmarshal req")
	}

	switch req.Opt {
	case skip:
		mediaSkip(req.Nick)
	case connect:
		m.Lock()
		clients = append(clients, req.Nick)
		m.Unlock()
	case leave:
		m.Lock()
		for i, nick := range clients {
			if req.Nick == nick {
				clients[i] = clients[len(clients)-1]
				clients[len(clients)-1] = ""
				clients = clients[:len(clients)-1]
			}
		}
	}
	_, err = conn.Write([]byte("msg recived"))
	if err != nil {
		log.Printf("Failed to write to client %s", conn.RemoteAddr())
	}
}

func mediaSkip(nick string) {
	log.Printf("%s asked to skip", nick)
}
