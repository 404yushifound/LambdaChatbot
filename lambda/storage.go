package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ChatEntry struct {
	Time string `json:"time"`
	User string `json:"user"`
	Bot  string `json:"bot"`
}

func SaveChat(userMsg, botReply string, t time.Time) error {
	date := t.Format("2006-01-02")
	fileName := fmt.Sprintf("chat_logs/%s.json", date)

	_ = os.MkdirAll("chat_logs", os.ModePerm)

	entry := ChatEntry{
		Time: t.Format("15:04:05"),
		User: userMsg,
		Bot:  botReply,
	}

	var history []ChatEntry

	// Load existing data
	data, err := os.ReadFile(fileName)
	if err == nil {
		_ = json.Unmarshal(data, &history)
	}

	history = append(history, entry)

	finalData, _ := json.MarshalIndent(history, "", "  ")
	return os.WriteFile(fileName, finalData, 0644)
}
