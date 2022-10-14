package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		line, _ := reader.ReadString('\n')
		fmt.Print(":: ", line)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter text: ")
		text, _ := stdin.ReadString('\n')
		if text == "/quit\n" {
			break
		} else {
			fmt.Fprintln(conn, text)
		}
	}
}

func main() {
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	conn, _ := net.Dial("tcp", *addrPtr)
	// Get the server address and port from the commandline arguments.
	go read(conn)
	write(conn)
	//TODO Start getting and sending user messages.
	//TODO Start asynchronously reading and displaying messages
}
