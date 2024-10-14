package handlers

import (
	"net"

	"github.com/charmbracelet/log"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func EchoHandler(conn net.Conn, arg string) {
	defer conn.Close()

	req, err := responses.ExtractRequest(conn)
	if err != nil {
		responses.NotFound(conn)
	}
	log.Info("EchoHandler ->", req.URL.Path)
	responses.RespondWithBody(arg, conn)
	return
}
