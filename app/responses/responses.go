package responses

import (
	"bufio"
	"net"
	"net/http"
)

var crlf = "\r\n"


func RespondOK(connection net.Conn) (bool, error) {
	ok_response := "HTTP/1.1 200 OK\r\n"

	_, err := connection.Write([]byte(ok_response + crlf))

	if err != nil {
		return false, err
	}
	return true, nil
}

func NotFound(connection net.Conn) (bool, error) {
	not_found_response := "HTTP/1.1 404 Not Found\r\n"

	_, err := connection.Write([]byte(not_found_response + crlf))
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractResponse(connection net.Conn) (http.Response, error) {
	reader := bufio.NewReader(connection)
	response, err := http.ReadResponse(reader, nil)

	if err != nil {
		return http.Response{}, err
	}

	return *response, nil
}
