package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/codecrafters-io/http-server-starter-go/app/responses"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	r := &utils.Router{}

	r.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("vou me matar hoje as 23:47"))
	})

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

			r.ServeHTTP(req, conn)
		}(conn)
	}
}
