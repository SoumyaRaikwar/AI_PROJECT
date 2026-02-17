package main

import (
	"fmt"
	"context"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
     
	llm, err := ollama.New(ollama.WithModel("llama3.2:latest"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	prompt := "What is the capital of France?"

	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		prompt,
	)
	if err != nil {
		log.Fatal(err)
	    }
	
	fmt.Println(completion)
}