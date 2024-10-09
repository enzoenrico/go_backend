package handlers

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func EchoHandler(conn net.Conn) {
	req, err := responses.ExtractRequest(conn)
	if err != nil{
		responses.NotFound(conn)
	}
	fmt.Println("EchoHandler ->" , req.URL.Path)

	responses.RespondOK(conn)
}