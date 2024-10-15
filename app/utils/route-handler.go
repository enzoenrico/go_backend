package utils

import (
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
			handlers.FileHandler(conn, split_path[len(split_path)-1])
		} else {
			responses.NotFound(conn)
		}
	default:
		responses.NotFound(conn)
	}
	return
}
