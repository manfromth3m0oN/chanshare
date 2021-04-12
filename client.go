package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func clientConnect(hostIp string) {
	addr := hostIp + ":42069"
	conn, _ := net.Dial("tcp", addr)
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)
	req := Req{}
	req.Nick = "m0on"
	req.Opt = connect
	enc.Encode(req)
	for {
		var resp Resp
		dec.Decode(&resp)
		if resp.Board != "" {
			log.Println(resp)
			mediaPos = resp.Pos
			loadThread(fmt.Sprintf("%d", resp.Thread), resp.Board)

		}
	}
}
