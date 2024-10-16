package handlers

import (
	"fmt"
	"net"
	"os"
)

func FileHandler(conn net.Conn, filename string) {
	defer conn.Close()
	dir := os.Args[2]

	fmt.Println("[+]Accessing the tmp directory:")
	data, err := os.ReadFile(dir + filename)
	if err != nil {
		return
	}
	conn.Write(data)
	return
}
