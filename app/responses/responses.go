package responses

import (
	"bufio"
	"net"
	"net/http"
)

var crlf = "\r\n"

// ConnectionHandler wraps the net.Conn allowing methods to operate on it
type ConnectionHandler struct {
	Connection net.Conn
}

// RespondOK sends a 200 OK response
func (ch *ConnectionHandler) RespondOK() (bool, error) {
	okResponse := "HTTP/1.1 200 OK\r\n"
	_, err := ch.Connection.Write([]byte(okResponse + crlf))
	if err != nil {
		return false, err
	}
	return true, nil
}

// NotFound sends a 404 Not Found response
func (ch *ConnectionHandler) NotFound() (bool, error) {
	notFoundResponse := "HTTP/1.1 404 Not Found\r\n"
	_, err := ch.Connection.Write([]byte(notFoundResponse + crlf))
	if err != nil {
		return false, err
	}
	return true, nil
}

// ExtractResponse reads the HTTP response from the connection
func (ch *ConnectionHandler) ExtractResponse() (http.Response, error) {
	reader := bufio.NewReader(ch.Connection)
	response, err := http.ReadResponse(reader, nil)

	if err != nil {
		return http.Response{}, err
	}

	return *response, nil
}
