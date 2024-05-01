package flagparser

import "flag"

func ParseFlag() (string, string) {
	model := flag.String("model", "llama3", "model to use")
	prompt := flag.String("prompt", "Could you tell me a joke?", "prompt to use")
	flag.Parse()

	return *model, *prompt
}
