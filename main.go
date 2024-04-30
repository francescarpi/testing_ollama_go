package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ollama-go/flagparser"
	"ollama-go/models"

	"github.com/fatih/color"
)

const apiHost = "http://127.0.0.1:11434/api"

func main() {
	model, prompt := flagparser.ParseFlag()

	client := &http.Client{}

	version, err := ollama_version(client)
	if err != nil {
		log.Fatal("The ollama server is down?")
		return
	}

	color.Green("Ollama version: %s\n", version)
	color.Green("Using model: %s\n", model)
	color.Blue("Waiting for the response...\n\n")

	var request models.Request = models.Request{
		Model:  model,
		Prompt: "Can you tell me a joke?",
		Stream: false,
	}

	if prompt != "" {
		request.Prompt = prompt
	}

	request_bytes, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, apiHost+"/generate", bytes.NewBuffer(request_bytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	var responseBody models.Response
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseBody.Response)
}

func ollama_version(client *http.Client) (string, error) {
	req, _ := http.NewRequest(http.MethodGet, apiHost+"/version", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	var versionResponse models.VersionResponse
	err = json.NewDecoder(resp.Body).Decode(&versionResponse)
	if err != nil {
		return "", err
	}

	return versionResponse.Version, nil
}
