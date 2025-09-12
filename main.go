package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go server | client <name>")
		return
	}

	switch os.Args[1] {
	case "server":
		server := NewServer()
		server.Run()
	case "client":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go client <username>")
			return
		}
		RunClient(os.Args[2])
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
