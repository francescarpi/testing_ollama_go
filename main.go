package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Request struct {
	Model    string  `json:"model"`
	Prompt   string  `json:"prompt"`
	System   *string `json:"system,omitempty"`
	Template *string `json:"template,omitempty"`
	Stream   bool    `json:"stream"`
	Options  struct {
		Temperature float64 `json:"temperature"`
	} `json:"options"`
}

type Response struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	Context            []int     `json:"context"`
	TotalDuration      int       `json:"total_duration"`
	LoadDuration       int       `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int       `json:"eval_duration"`
}

func main() {
	var request Request = Request{
		Model:  "llama3",
		Prompt: "Can you tell me a joke?",
		Stream: false,
	}

	request_bytes, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:11434/api/generate", bytes.NewBuffer(request_bytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	var responseBody Response
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseBody.Response)
}
