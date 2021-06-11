package main

import (
	"log"
	"time"
)

func main() {
	t := getTheTime()

	log.Println("Current Time:", t)
}

func getTheTime() string {
	t := time.Now().UTC()
	return t.Format(time.UnixDate)
}
