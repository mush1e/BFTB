package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s - usage: \n", os.Args[0])
		log.Fatalf("Expected exactly one argument; got %d", len(os.Args)-1)
	}

	host := os.Args[1]
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("lookup ip : %q | err : %q", host, err)
	}
	if len(ips) == 0 {
		log.Fatalf("could not find any ips for host : %q", host)
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			fmt.Println(ip)
			goto IPV6
		}
	}
	fmt.Println("ipv4 - none")

IPV6:
	for _, ip := range ips {
		if ip.To4() == nil {
			fmt.Println(ip)
			return
		}
	}
	fmt.Println("ipv6 - none")
}
