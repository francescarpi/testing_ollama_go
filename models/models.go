package models

import (
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

type VersionResponse struct {
	Version string `json:"version"`
}
