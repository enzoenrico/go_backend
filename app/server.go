package main

import (
	"fmt"
	"net"
	"os"

	"github.com/charmbracelet/log"
	"github.com/codecrafters-io/http-server-starter-go/app/handlers"
	"github.com/codecrafters-io/http-server-starter-go/app/responses"
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
			log.Errorf("error accepting connection: %s", err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			req, err := responses.ExtractRequest(conn)

			if err != nil {
				log.Errorf("error accepting connection: %s", err)
			}

			log.Logf(0,"> %s request on %s", req.Method, req.URL.Path)

			handlers.EchoHandler(conn, "oki")

		}(conn)
	}
}
