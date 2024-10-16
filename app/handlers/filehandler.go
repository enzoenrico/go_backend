package handlers

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func FileHandler(conn net.Conn, filename string) {
	defer conn.Close()
	dir := os.Args[2]

	fmt.Println("[+]Accessing the tmp directory:")
	data, err := os.ReadFile(dir + filename)
	if err != nil {
        responses.NotFound(conn)
		return
	}
    responses.RespondWithFile(data, conn)
	return
}
