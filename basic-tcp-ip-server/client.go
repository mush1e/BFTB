package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Client(port int) {
	log.SetPrefix("client\t")

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()
	go func() {
		for scanner := bufio.NewScanner(conn); scanner.Scan(); {
			fmt.Printf("%s\n", scanner.Text()) // note: printf doesn't add a newline, so we need to add it ourselves

			if err := scanner.Err(); err != nil {
				log.Fatalf("error reading from %s: %v", conn.RemoteAddr(), err)
			}
			if scanner.Err() != nil {
				log.Fatalf("error reading from %s: %v", conn.RemoteAddr(), err)
			}
		}
	}()

	for stdinScanner := bufio.NewScanner(os.Stdin); stdinScanner.Scan(); { // find the next newline in stdin
		log.Printf("sent: %s\n", stdinScanner.Text())
		if _, err := conn.Write(stdinScanner.Bytes()); err != nil { // scanner.Bytes() returns a slice of bytes up to but not including the next newline
			log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
		}
		if _, err := conn.Write([]byte("\n")); err != nil { // we need to add the newline back in
			log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
		}
		if stdinScanner.Err() != nil {
			log.Fatalf("error reading from %s: %v", conn.RemoteAddr(), err)
		}
	}
}
