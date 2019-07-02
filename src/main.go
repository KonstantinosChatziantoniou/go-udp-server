package main

import (
	"dbcontroller"
	"log"
	"net"
)

const (
	// PORT the port where we will listen for udp packets
	PORT = "8080"
)

func main() {

	// Setup MYSQL connection
	dbcontroller.InitializeMYSQL()

	// Setup UDP listener
	pc, err := net.ListenPacket("udp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}

	defer pc.Close()

	// Listening for packets
	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		dbcontroller.Serve(buf, addr.String(), n)
	}
}
