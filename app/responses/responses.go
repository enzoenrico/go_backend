package responses

import (
	"net"
)

func RespondOK(connection net.Conn) bool {
	_, err := connection.Write([]byte("HTTP/1.1 200 OK\r\n"))
	if err != nil{
		return false
	}
	return true
}
