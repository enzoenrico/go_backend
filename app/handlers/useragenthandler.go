package handlers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func UserAgentHandler(conn net.Conn, req http.Request) string {
	defer conn.Close()
	fmt.Println(req.UserAgent())

	responses.RespondWithBody("userAgent", conn)

	return req.UserAgent()
}
