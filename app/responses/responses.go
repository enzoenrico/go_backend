package responses

import (
	"net"
)

func RespondOK(connection net.Conn) bool {
	ok_response := "HTTP/1.1 200 OK\r\n"

	crlf := "\r\n"

	_, err := connection.Write([]byte(ok_response + crlf))

	if err != nil {
		return false
	}
	return true
}
