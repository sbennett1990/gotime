package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"

	"suah.dev/protect"
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

	privDrop()

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

// Drop privileges.
// This should be called even in debug mode.
func privDrop() {
	if !debug {
		pw, err := user.Lookup(_PW_USER)
		if err != nil {
			log.Fatal(err)
		}
		err = syscall.Chroot(_PW_DIR)
		if err != nil {
			log.Fatal(err)
		}
		err = syscall.Chdir("/")
		if err != nil {
			log.Fatal(err)
		}

		uid, err := strconv.Atoi(pw.Uid)
		if err != nil {
			log.Fatal(err)
		}

		gid, err := strconv.Atoi(pw.Gid)
		if err != nil {
			log.Fatal(err)
		}

		err = syscall.Setgroups([]int{gid})
		if err != nil {
			log.Fatal(err)
		}

		err = syscall.Setregid(gid, gid)
		if err != nil {
			log.Fatal(err)
		}

		err = syscall.Setreuid(uid, uid)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := protect.Pledge("stdio inet proc")
	if err != nil {
		log.Fatalln("pledge failed")
	}
}
