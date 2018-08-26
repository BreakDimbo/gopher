package main

import (
	"flag"
	"net"
)

var port = flag.Int("port", 8000, "tcp port to listen")

func main() {
	flag.Parse()

}

func handleConn(c net.Conn) {

}
