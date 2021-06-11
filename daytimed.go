package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

//	"github.com/poolpOrg/ipcmsg"
//	"github.com/poolpOrg/privsep"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "13"
	CONN_TYPE = "tcp"
)

func main() {
	// Need to be root to bind to low port...

	ln, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		fmt.Println("error listening: ", err.Error())
		os.Exit(1)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 256)

	conn.Write([]byte(""))
	conn.Close()
}

func getTheTime() string {
	t := time.Now().UTC()
	return t.Format(time.UnixDate)
}
