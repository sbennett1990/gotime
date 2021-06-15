/*
 * Copyright (c) 2021 Scott Bennett <scottb@fastmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

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

//
// Handle incoming connections.
//
func handleRequest(conn net.Conn) {
	timestr := getTheTime()
	conn.Write([]byte(timestr))
	conn.Close()
}

//
// Return the current UTC time as a human-readable string.
//
func getTheTime() string {
	t := time.Now().UTC()
	return t.Format(time.UnixDate)
}

//
// Drop privileges.
// This should be called even in debug mode.
//
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
