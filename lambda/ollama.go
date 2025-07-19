package main

import (
	"github.com/go-resty/resty/v2"
	"encoding/json"
) 

type OllamaRequest struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func GenerateReply(prompt string) (string, error) {
	client := resty.New()

	body := OllamaRequest{
		Model: "llama3",
		Prompt: prompt,
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("http://localhost:11434/api/generate")

	if err != nil {
		return "", err
	}

	var res OllamaResponse
	if err := json.Unmarshal(resp.Body(), &res); err != nil {
		return "", err
	}

	return res.Response, nil
}
