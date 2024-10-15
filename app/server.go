package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/codecrafters-io/http-server-starter-go/app/responses"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
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
			// defer func() {
			// 	if err := conn.Close(); err != nil {
			// 		log.Errorf("error closing connection: %s", err)
			// 	} else {
			// 		fmt.Println("> Connection closed.")
			// 	}
			// }()
			defer conn.Close()

			req, err := responses.ExtractRequest(conn)

			if err != nil {
				log.Errorf("error accepting connection: %s", err)
			}

			fmt.Printf("> %s request on %s\n", req.Method, req.URL.Path)

			split_path := strings.Split(strings.Trim(req.URL.Path, "/"), "/")

			fmt.Printf("\t> Split path: %s \n", split_path)

			// for _, v := range split_path {

			// 	fmt.Printf("\t> %s \n", v)
			// }

			utils.RouteHandler(conn, split_path, req)

			fmt.Println("> Response sent.")
		}(conn)
	}
}
