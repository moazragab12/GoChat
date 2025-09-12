package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func RunClient(username string) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", username)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				break
			}
			msg, err := (&Message{}).ToString(line)
			if err == nil {
				fmt.Printf("[%s]: %s\n", msg.From, msg.Content)
			}
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome,", username, "! Type /quit to exit.")
	fmt.Println("Use '@bob:Hello' for private messages")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "/quit" {
			return
		}

		msg := &Message{From: username, Content: text}
		if strings.HasPrefix(text, "@") {
			parts := strings.SplitN(text, ":", 2)
			if len(parts) == 2 {
				msg.To = strings.TrimPrefix(parts[0], "@")
				msg.Content = parts[1]
			}
		}

		conn.Write(msg.ToJSON())
	}

}
