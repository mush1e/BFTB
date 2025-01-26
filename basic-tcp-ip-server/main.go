package main

import "flag"

func main() {
	mode := flag.String("m", "server", "set mode to client or server")
	port := flag.Int("p", 8080, "set appropriate port to connect to / host on")
	flag.Parse()

	switch *mode {
	case "server":
		Server(*port)
	case "client":
		Client(*port)
	}
}
