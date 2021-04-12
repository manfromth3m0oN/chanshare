package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func clientConnect(hostIp string) {
	addr := hostIp + ":42069"
	conn, _ := net.Dial("tcp", addr)
	req := Req{}
	req.Nick = "m0on"
	req.Opt = connect
	msg, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Couldnt marshal request: %v", err)
	}
	fmt.Fprint(conn, msg)
}
