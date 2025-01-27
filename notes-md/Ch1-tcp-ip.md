#### **What is the internet?**
- The internet is a network of inter-connected computers, where computers in the network can communicate with each other reliably
- So the internet solves 2 major problems 
	- **Routing** (IP)      
		- How can I send a message to a computer when im not directly connected to it
	- **Coherence** (TCP)
		- How do I make sure the information passed is in one piece in the right order
- Together both these concepts are called `TCP/IP` which is how the internet works

#### **So how does TCP actually work?**
- Essentially when communicating with a computer, you send packets of data 
- Each packet has a sequence number (which packet is this?) and a checksum (is this packet corrupted?)
- after receiving each packet the remote computer sends back an acknowledgement (like hey! I got packet 6 looks good! wohoooo!!) 
- If you don't get an acknowledgement, you resend the packet.
- If a packet is received with errors, it is discarded, and the sender eventually retransmits it.
- This whole back and forth ensures all the information gets through in order


#### **Okay now how does IP actually work?**
- IP is a bit more complicated so were going to abstract it. essentially this explanation is inaccurate but you'll get the idea
- Each computer has an `address`, which is an identifier that tells other computers how to get to it, this address is called an `IP Address`
- They also have a list of known networks and the best next-hop addresses to reach them, contained in a `routing table`
- When you try to send a message to another computer, your computer looks at the `Routing Table` to see if it knows how to get to that computer
	- if it does, it sends the message to the next computer in the chain 
	- otherwise, it sends the message to the next computer that might know how to get to the target
- you kinda keep repeating this process till you get to the target computer
- if theres no path from your computer to the other computer, your message fails 


#### **Going deeper into addresses, and what the heck is a port?**
- Okay now that we've talked about sending this "message", how do we actually send this message
- well we need 2 things to do this
	- the `IP Address` of the computer we want to send the message to
	- the `Port` the target computer is listening on
- The **port** is essentially a "door" on the computer, and the **IP address** tells the network where the computer is located. You need both to send your message, i.e after going to the address you need to know which door to knock on. 


#### **IP Address Types: What Are They?**
- IP Addresses come in 2 forms 
	- `IPv4` a 32 bit number which could look like `DDD.DDD.DDD.DDD`, where DDD is a number between 0 and 255.
	- `IPv6` a 128 bit number which could look like `XXXX:XXXX:XXXX:XXXX:XXXX:XXXX:XXXX:XXXX`, where XXXX is a 16-bit hexadecimal number; that is, each X is one of `0..=9` or `a..=f`

| IP Address                              | Type | Note                                                                                                     |
| --------------------------------------- | ---- | -------------------------------------------------------------------------------------------------------- |
| 192.168.000.001                         | IPv4 | private network IP address refers to hosting computer                                                    |
| 192.168.0.1                             | IPv4 | Same as above; you can omit leading zeroes                                                               |
| 0000:0000:0000:0000:0000:ffff:c0a8:0001 | IPv6 | Refers to the same computer as above; IPv4 addresses can be embedded in IPv6 by prefixing with `::ffff:` |
| ::ffff:c0a8:0001                        | IPv6 | Same as above; you can omit leading zeroes                                                               |


#### Understanding Ports: What Are They and Why Do They Matter?
- Its pretty common for one computer to want to host multiple web services that behave in different ways
- for example you could want to host a `Minecraft` server and also some `todo list` project at the same time
- since they both are on the same computer, they have the same address, so how do you distinguish if someone is trying to join your Minecraft server or is just trying to add grocery shopping to their todo list.
- Say hello to **Ports** :D, to deal with this issue you can assign a Port to each of the services you wanna host.
- In english well you can specify which door the "message" has to knock on to get access to the service they are looking for. like Minecraft could be behind door 102 and todo list behind 103.  
- PS **`a Port is just a number between 0-65535`**
- you cant just assign any port for any service you'd like btw. Some ports have special duties but this is out of our scope for now.


#### **Now lets implement something to demonstrate TCP/IP**
- Okay so lets write a dumb program to kinda solidify our understanding of how to do this in go
- first lets outline what were going to build! Were going to build a TCP/IP client and server
- what the client does is send the server some text and gets back the same text capitalized
- Simple enough right? lets get started. 
- Now our code can run in 2 modes, client and server. They also probably need some port to send messages to and listen on respectively so lets kinda tackle this in our `main.go` file

```go
func main() {
	mode := flag.String("m", "server", "set the mode to [server] or [client]")
	port := flag.Int("p", 8080, "set the port to listen or query")
	flag.Parse()

	switch *mode {
	case "server":
		Server(*port)
	case "client":
		Client(*port)
	}
}
```

- `flag`'s  are just ways of taking command line arguments in go, so when you try to run our program you can specify what `mode` and `port` you run on / use

```bash
./basic-tcp-ip-example -m server -p 8080  # To run in server mode
./basic-tcp-ip-example -p 8080 -m client  # To run in client mode
```

#### **Tackling the Server implementation first**

- so first lets write a function that takes the request coming in and sends back the text capitalized

```go
func generateResponse(writer io.Writer, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text() 
		fmt.Fprintf(writer, "%v\n", strings.ToUpper(line))
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}
}
```
- lets break down this function so we can make sense of what's going on! 
- we're defining a function called `generateResponse` that takes in a `io.Reader` and `io.Writer` interface
- essentially any type that has a `Read` and `Write` function associated with it
- We keep reading lines from `reader` and then capitalizing the line read and sending it back to `writer`

- Alright lets write a function to actually handle the server stuff (what we're calling in our `main.go`)

```go
func Server(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("error listening on port %d : %q", port, err)
	} 
	defer listener.Close() 
	log.Printf("server listening at localhost on port : %d", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		
		go func(c net.Conn) { 
			defer c.Close() 
			generateResponse(c, c) 
		}(conn)
	}
}
```

- essentially we create try listening on `port` over `localhost`, handle errors if it doesnt work out, then just have a loop that keeps polling (fancy word for waiting) for connection requests.
- once we do accept a connection and everything is okay we just spawn a `goroutine` to generate and send our response via `getResponse(conn, conn)` (`net.Conn` satisfies `io.Reader` and `io.Writer` since it implements a `Read` and `Write` function)

#### **Onto the client**
- Now the client isnt going to exactly be as straight forward as the server since essentially we're doing 2 separate things! 
	- Sending out messages to the server
	- Receiving responses from the server
- Now we don't want either of these tasks to block each other, so again were going to rely on `goroutines`
- Lets start implementing our client function

```go
func Client(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("unable to establish connection with server at : %d", port)
	}
	defer conn.Close() 
	// listening and sending implementation
}
```

- okay so far so good? we attempt to establish a connection to the server, do some error handling and just clean up conn after we're done 
- Now on to the juicy stuff! lets start implementing our listener in a `goroutine` to listen for any response messages coming from the server and print them to `stdout`

```go
// Within our client function
go func() {
	for scanner := bufio.NewScanner(conn); scanner.Scan(); {
		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading from %s: %v", conn.RemoteAddr(), err)
		}
		fmt.Printf("recieved from server - %q\n", scanner.Text())
	}
}()
```

- well we just kinda keep looping and waiting for our scanner to read something coming in over `conn`, once we get something, we just display it to the terminal
- now for the sending part

```go
for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
	fmt.Printf("sending over %q to the server\n", scanner.Text())
	msg := scanner.Text() + '\n'
	
	if _, err := conn.Write([]byte(msg)); err != nil {
		log.Printf("error writing to server: %v\n", err) 
		 break 
	}
}
if err := scanner.Err(); err != nil { 
	log.Fatalf("error reading from stdin: %v\n", err) 
}

```

