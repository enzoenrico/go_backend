package handlers

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func EchoHandler(conn net.Conn, arg string) {
	defer conn.Close()

	fmt.Println("> Entering the Echo Handler")

	req, err := responses.ExtractRequest(conn)
	if err != nil {
		responses.NotFound(conn)
	}
	fmt.Println("EchoHandler ->", req.URL.Path)
	responses.RespondWithBody(arg, conn)
	return
}
