package main

import (
	"fmt"
	"context"
	"log"

	"github.com/tmc/langchaingo/llms"
	OpenAI "github.com/tmc/langchaingo/llms/openai"
)

func main() {
     
	llm, err := OpenAI.New()
	if err != nil {
		log.Fatal(err)
	}

	prompt := "What year did Shahrukh Khan born?"

	completion, err := llms.GenerateFromSinglePrompt(
		context.Background(),
		llm,
		prompt,
	)
	if err != nil {
		log.Fatal(err)
	    }
	
	fmt.Println(completion)
}