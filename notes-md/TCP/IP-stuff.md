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
- after receiving each packet the remote computer sends back an acknowledgement (like hey! I got packet 6 looks good!) 
- If you don't get an acknowledgement, you resend the packet. Also if you get a corrupt packet acknowledgement you also resend it
- This whole back and forth ensures all the information gets through in order


#### **Okay now how does IP actually work?**
- IP is a bit more complicated so were going to abstract it. essentially this explanation is inaccurate but you'll get the idea
- Each computer has an `address`, which is an identifier that tells other computers how to get to it, this address is called an `IP Address`
- They also have a list of other known computers on them. this list is called a `Routing Table`
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

| IP Address                     | Type  | Note                                                             |
|---------------------------------|-------|------------------------------------------------------------------|
| 192.168.000.001                 | IPv4  | Localhost; refers to hosting computer                            |
| 192.168.0.1                     | IPv4  | Same as above; you can omit leading zeroes                       |
| 0000:0000:0000:0000:0000:ffff:c0a8:0001 | IPv6  | Refers to the same computer as above; IPv4 addresses can be embedded in IPv6 by prefixing with `::ffff:` |
| ::ffff:c0a8:0001                | IPv6  | Same as above; you can omit leading zeroes                       |


#### Understanding Ports: What Are They and Why Do They Matter?
- Its pretty common for one computer to want to host multiple web services that behave in different ways
- for example you could want to host a `Minecraft` server and also some `todo list` project at the same time
- since they both are on the same computer, they have the same address, so how do you distinguish if someone is trying to join your Minecraft server or is just trying to add grocery shopping to their todo list.
- Say hello to **Ports** :D, to deal with this issue you can assign a Port to each of the services you wanna host.
- In english well you can specify which door the "message" has to knock on to get access to the service they are looking for. like Minecraft could be behind door 102 and todo list behind 103.  
- PS **`a Port is just a number between 0-65535`**
- you cant just assign any port for any service you'd like btw. Some ports have special duties but this is out of our scope for now.

| Service   | Port |
|-----------|------|
| HTTP      | 80   |
| HTTPS     | 443  |
| SSH       | 22   |
| SMTP      | 25   |
| DNS       | 53   |
| FTP       | 21   |
| Postgres  | 5432 |
