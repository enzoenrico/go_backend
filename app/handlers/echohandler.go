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
	fmt.Println(req.URL.Path)
	if err != nil {
		fmt.Println("> Request could not be converted! \n")
		responses.NotFound(conn)
	}

	fmt.Println("> EchoHandler -> ", req.URL.Path)
	responses.RespondWithBody(arg, conn)
	return
}
