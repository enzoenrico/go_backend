package responses

import (
	"bufio"
	"net"
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
)

var crlf = "\r\n"
var ok_response = "HTTP/1.1 200 OK\r\n"
var not_found_response = "HTTP/1.1 404 Not Found\r\n"

func RespondOK(connection net.Conn) (bool, error) {
	_, err := connection.Write([]byte(ok_response + crlf))

	if err != nil {
		return false, err
	}
	return true, nil
}

func NotFound(connection net.Conn) (bool, error) {
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
	response := boiler + body_len + crlf + crlf + body_content

	log.Infof("> Sending: %s", response)

	_, err := connection.Write([]byte(response))
	if err != nil {
		return false, err
	}
	return true, nil
}

func RespondWithFile(content []byte, connection net.Conn) (bool, error) {
	boiler := "HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: "
	contentLen := strconv.Itoa(len(content))
	response := boiler + contentLen + crlf + crlf

	// Log information about the response
	log.Infof("Sending binary data: Content-Length: %s", contentLen)

	// Send the response headers
	_, err := connection.Write([]byte(response))
	if err != nil {
		return false, err
	}

	// Send the actual file content
	_, err = connection.Write(content)
	if err != nil {
		return false, err
	}

	return true, nil
}
