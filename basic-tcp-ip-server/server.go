package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func generateResponse(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(writer, "%v\n", strings.ToUpper(line))
	}
	if scanner.Err() != nil {
		log.Fatalf(scanner.Err().Error())
	}
}

func Server(port int) {
	log.SetPrefix("server\t")

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer listener.Close()
	log.Printf("listening at localhost: %s", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go generateResponse(conn, conn)
	}
}
