package handlers

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

func FileHandler(conn net.Conn, filename string) {
	defer conn.Close()
	// TODO: implement file handling
	tmp_dir, err := os.ReadDir("/tmp")
	if err != nil {
		return false
	}

	fmt.Println("[+]Accessing the tmp directory:")
	for _, file := range tmp_dir {
		fmt.Printf("> /tmp/%s \n", file.Name())
		if file.Name() == filename {
			fmt.Printf("[!]File found! \n \r %s \n", file.Name())
		}
	}
	return 
}
