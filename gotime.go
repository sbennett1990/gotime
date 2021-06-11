package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"

//	"github.com/poolpOrg/ipcmsg"
//	"github.com/poolpOrg/privsep"
)

const (
	DEBUG_ADDR = "localhost"
	PORT       = "13"
	DEBUG_PORT = "13013"
	_PW_USER   = "_daytimed"
	_PW_DIR    = "/var/empty"
)

var debug bool

func main() {
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.Parse()
	if flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Need to be root to bind to low port
	if !debug {
		if os.Geteuid() != 0 {
			log.Fatalln("need root privileges")
		}
	}

	listenAddr := ":" + PORT
	if debug {
		listenAddr = DEBUG_ADDR + ":" + DEBUG_PORT
	}

	ln, err := net.Listen("tcp4", listenAddr)
	if err != nil {
		log.Fatalln("error listening on addr", listenAddr, "error:", err.Error())
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("error accepting: ", err.Error())
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
