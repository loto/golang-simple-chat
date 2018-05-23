package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

// RunHost takes an ip as argument and listens for messages
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}

	fmt.Println("Listening on ", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}

	fmt.Println("New connection accepted")

	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Println("Message received: ", message)

	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyReadErr := replyReader.ReadString('\n')
	if replyReadErr != nil {
		log.Fatal("Error: ", replyReadErr)
	}
	fmt.Fprint(conn, replyMessage)
}

// RunGuest takes an ip as argument and blabla
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Error: ", dialErr)
	}

	for {
		handleGuest(conn)
	}
}

func handleGuest(conn net.Conn) {
	for {
		fmt.Print("Send message: ")
		reader := bufio.NewReader(os.Stdin)
		message, readErr := reader.ReadString('\n')
		if readErr != nil {
			log.Fatal("Error: ", readErr)
		}
		fmt.Fprint(conn, message)

		replyReader := bufio.NewReader(conn)
		replyMessage, replyReadErr := replyReader.ReadString('\n')
		if replyReadErr != nil {
			log.Fatal("Error: ", replyReadErr)
		}
		fmt.Println("Message received: ", replyMessage)
	}
}
