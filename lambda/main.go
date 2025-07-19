package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("🧠 Welcome to Lambda Chatbot (type 'exit' to quit)")

	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" {
			break
		}
		if input == "" {
			continue
		}

		reply, err := GenerateReply(input)
		if err != nil {
			fmt.Println("Bot: (error generating response)", err)
			continue
		}

		fmt.Println("Bot:", reply)

		err = SaveChat(input, reply, time.Now())
		if err != nil {
			fmt.Println("⚠️ Failed to save chat:", err)
		}
	}
}
