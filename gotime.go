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

// Handle incoming connections.
func handleRequest(conn net.Conn) {
	// make a buffer to handle incoming data
	//buf := make([]byte, 256)

	timestr := getTheTime()
	conn.Write([]byte(timestr))
	conn.Close()
}

// Return the current UTC time as a human-readable string.
func getTheTime() string {
	t := time.Now().UTC()
	return t.Format(time.UnixDate)
}
