package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	defer c.Close()
}

func echo(c net.Conn, s string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}
