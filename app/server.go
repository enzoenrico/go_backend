package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/handlers"
	"github.com/codecrafters-io/http-server-starter-go/app/responses"

	"github.com/charmbracelet/log"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Errorf("Error accepting connection: %s", err)
			continue
		}
		go func(conn net.Conn) {
			log.Info("New conn from: ", conn.LocalAddr().String())
			defer conn.Close()
			req, err := responses.ExtractRequest(conn)
			if err != nil {
				responses.NotFound(conn)
				return
			}
			// gotta get the path fully and it's arguments
			split_path := strings.Split(req.URL.Path, "/")
			
			if len(split_path) > 2{
				handlers.EchoHandler(conn, split_path[2])
			}
		}(conn)
	}
}
