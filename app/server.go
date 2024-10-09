package main

import (
	"fmt"
	"net"
	"os"
	"responses"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	responses.RespondOK(conn)

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
