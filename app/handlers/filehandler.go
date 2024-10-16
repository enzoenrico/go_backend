package handlers

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func GetFileHandler(conn net.Conn, filename string) {
	defer conn.Close()
	dir := os.Args[2]

	fmt.Println("[+]Accessing the tmp directory:")
	data, err := os.ReadFile(dir + filename)
	if err != nil {
		responses.NotFound(conn)
		return
	}
	responses.RespondWithFile(data, conn)
	return
}

func PostFileHandler(conn net.Conn, filename string, body []byte) {
	defer conn.Close()
	dir := os.Args[2]

	//acc create the file on system
	fmt.Printf("> Creating: %s \n", dir+filename)
	fmt.Printf("\t Body: %s \n", string(body))

	err := os.WriteFile(dir+filename, body, 0644)
	if err != nil {
		responses.NotFound(conn)
	}

	crlf := "\r\n"
	response := "HTTP/1.1 201 Created\r\nContent-Length:"

	conn.Write([]byte(response + strconv.Itoa(len(body)) + crlf + crlf))
}
