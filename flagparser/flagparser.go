package flagparser

import "flag"

func ParseFlag() (string, string) {
	model := flag.String("model", "llama3", "model to use")
	prompt := flag.String("prompt", "", "prompt to use")
	flag.Parse()

	return *model, *prompt
}
