package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	l, err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	_, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
