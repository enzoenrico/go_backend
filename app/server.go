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
	defer conn.Close()

	if err != nil{
		panic(err)
	}

	resp, err := responses.ExtractResponse(conn)
	if err != nil{
		fmt.Println(resp)
		panic(err)
	}
	path := utils.ExtractPath(resp)

	if ok := utils.ValidPath(path); ok{
		responses.RespondOK(conn)
	}else{
		responses.NotFound(conn)
	}

}
