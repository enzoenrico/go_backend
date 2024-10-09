package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()

	con_handler := responses.ConnectionHandler{Connection: conn}

	res, err := con_handler.ExtractResponse()
	if err != nil {
		fmt.Println("Error getting path: ", err.Error())
	}
	path := res.Request.URL.Path
	if ok := utils.ValidPath(path); ok{
		con_handler.RespondOK()
	}else{
		con_handler.NotFound()
	}
}
