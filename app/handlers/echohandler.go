package handlers

import (
	"github.com/charmbracelet/log"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func EchoHandler(conn net.Conn, arg string) {
	req, err := responses.ExtractRequest(conn)
	if err != nil {
		responses.NotFound(conn)
	}
	log.Info("EchoHandler ->", req.URL.Path)
	// responses.RespondOK(conn)
	responses.RespondWithBody(arg, conn)
}
