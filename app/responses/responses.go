package responses

import (
	"bufio"
	"net"
	"net/http"
	"strconv"
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

func ExtractRequest(connection net.Conn) (*http.Request, error) {
	reader := bufio.NewReader(connection)

	response, err := http.ReadRequest(reader)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func RespondWithBody(body_content string, connection net.Conn) (bool, error) {
	boiler := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: "
	body_len := strconv.Itoa(len(body_content))
	response := boiler + body_len + crlf + body_content + crlf

	_, err := connection.Write([]byte(response))
	if err != nil {
		return false, err
	}
	return true, nil
}
