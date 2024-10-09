package responses

import (
	"bufio"
	"fmt"
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

func ExtractResponse(connection net.Conn) (*http.Response, error) {
    // Create a buffered reader from the connection.
    reader := bufio.NewReader(connection)
    
    // Use http.ReadResponse to parse and get the HTTP response.
    response, err := http.ReadResponse(reader, nil)
    if err != nil {
        return nil, err
    }
    
    return response, nil
}
