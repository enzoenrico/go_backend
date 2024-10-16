package utils

import (
	"fmt"
	"net"
	"net/http"

	"github.com/codecrafters-io/http-server-starter-go/app/handlers"
	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func RouteHandler(conn net.Conn, split_path []string, request *http.Request) {
	switch split_path[0] {
	case "":
		responses.RespondOK(conn)

	case "echo":
		handlers.EchoHandler(conn, split_path[len(split_path)-1])

	case "user-agent":
		handlers.UserAgentHandler(conn, *request)

	case "files":
		if len(split_path) >= 1 {
			if request.Method == "GET" {
				handlers.GetFileHandler(conn, split_path[len(split_path)-1])
			}
			if request.Method == "POST" {
				fmt.Printf("> Body content: \n\t %s \n", request.Body)
				// TODO: Read body content
				handlers.PostFileHandler(conn, split_path[len(split_path)-1], []byte("penis"))
			}
		} else {
			responses.NotFound(conn)
		}
	default:
		responses.NotFound(conn)
	}
	return
}
