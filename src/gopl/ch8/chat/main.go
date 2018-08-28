package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message
var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		case msg := <-message:
			for ch := range clients {
				ch <- msg // use bufferd ch and non-block send 
			}
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)   // outgoing message
	go clientWriter(conn, ch) // send message to client

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	message <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ": " + input.Text()
	}

	leaving <- ch
	message <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for content := range ch {
		fmt.Fprintln(conn, content)
	}
}
