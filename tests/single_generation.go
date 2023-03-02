package main

import (
    "context"
    "log"
    "github.com/honeyhive-ai/honeyhive-go-sdk"
    "github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
    "github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/operations"
	"fmt"
)

func main() {
    s := honeyhive.New(honeyhive.WithSecurity(
        shared.Security{
            BearerAuth: shared.SchemeBearerAuth{
                Authorization: "Bearer bRyCnbMZRMiVOMVx41TSvJr8ZHbW3qDG",
            },
        },
    ))
    
	taskName := "Sandbox - Email Writer"
	prompt := "Say this is a {{var}}"
	model := "text-davinci-003"
	inputs := map[string]interface{}{
		"topic": "test",
		"tone": "friendly",
	}

	response := map[string]interface{}{
		"id": "cmpl-6oIrGYG5V1GJi57QhsoGWOs1AHwb7",
		"object": "text_completion",
		"created": 1677446910,
		"model": "text-davinci-003",
		"choices": []map[string]interface{}{
			{
				"text": "\n\nThis is indeed a test",
				"index": 0,
				"logprobs": nil,
				"finish_reason": "length",
			},
		},
		"usage": map[string]interface{}{
			"prompt_tokens": 5,
			"completion_tokens": 7,
			"total_tokens": 12,
		},
	}
	source := "curie-writer"
	temp := 0.5
	topP := 1.0
	freqPenalty := 0.0
	presPenalty := 0.0
	max_tokens := 100
	hyperparameters := shared.HyperParameters{
		FrequencyPenalty: &freqPenalty,
		MaxTokens: &max_tokens,
		PresencePenalty: &presPenalty,
		Temperature: &temp,
		TopP: &topP,
	}
	generation := "\n\nThis is indeed a test"
	latency := 1000.23


    
    req := operations.IngestSingleGenerationRequest{
        Request: shared.SingleGenerationInput{
            Task: &taskName,
			Prompt: &prompt,
			Model: &model,
			Inputs: inputs,
			Response: response,
			Source: &source,
			Hyperparameters: &hyperparameters,
			Generation: &generation,
			Latency: &latency,
        },
    }

    ctx := context.Background()
    res, err := s.Generation.IngestSingleGeneration(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SingleGenerationOutput != nil {
        // handle response
		fmt.Println(*res.SingleGenerationOutput.GenerationID)
    }
}