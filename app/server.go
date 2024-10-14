package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/codecrafters-io/http-server-starter-go/app/handlers"
	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func main() {
	fmt.Println(">Program started on http://localhost:4221/")
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

			fmt.Printf("> %s request on %s\n", req.Method, req.URL.Path)

			split_path := strings.Split(req.URL.Path, "/")
			fmt.Printf("> Split path: %s \n", split_path)
			fmt.Printf("%T", split_path)

			switch split_path[0] {
			case "":
				responses.RespondOK(conn)
			case "/echo":
				handlers.EchoHandler(conn, split_path[1])
				// conn.Write([]byte("vou me matar"))
			}

			fmt.Println("> Response sent.")
			fmt.Println("> Program finished.")

		}(conn)
	}
}
