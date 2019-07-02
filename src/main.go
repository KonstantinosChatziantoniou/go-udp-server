package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	// PORT the port where we will listen for udp packets
	PORT = "8080"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)

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

		log.Println("len ", n, "addr ", addr.String(), " msg ", string(buf))
	}
}
