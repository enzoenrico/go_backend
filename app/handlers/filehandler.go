package handlers

import (
	"fmt"
	"net"
	"os"
)

func FileHandler(conn net.Conn, filename string) bool {
	defer conn.Close()
	// TODO: implement file handling
	tmp_dir, err := os.ReadDir("/tmp")
	if err != nil {
		return false
	}

	for _, file := range tmp_dir {
		fmt.Printf("> /tmp/%s \n", file.Name())
	}
	return true
}
