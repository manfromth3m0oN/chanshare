package main

import (
	"encoding/gob"
	"log"
	"net"
	"sync"
)

var clients []string

const (
	skip = iota
	connect
	leave
)

type Req struct {
	Nick string `json:"nick"`
	Opt  int    `json:"opt"`
}

type Resp struct {
	Board string
	Thread uint32
	Pos int
}

func hostFunc() {
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
	dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)

	var req Req
	dec.Decode(&req)

	log.Println(req)

	switch req.Opt {
	case skip:
		mediaSkip(req.Nick)
	case connect:
		m.Lock()
		clients = append(clients, req.Nick)
		m.Unlock()
		log.Println(clients)
		resp := Resp{
			Board: board,
			Thread: thread,
			Pos: mediaPos,
		}
		enc.Encode(resp)
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
}

func mediaSkip(nick string) {
	log.Printf("%s asked to skip", nick)
}
