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
	input := map[string]interface{}{
		"topic": "test",
		"tone": "friendly",
	}
	prompts := []string{
		"curie-writer",
	}

    req := operations.TaskCreateGenerationRequest{
        Request: shared.TaskGenerationInput{
            Task: &taskName,
			Prompts: prompts,
            Input: input,
        },
    }

    ctx := context.Background()
    res, err := s.Generation.TaskCreateGeneration(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskGenerationOutput != nil {
        // handle response
		fmt.Println(*res.TaskGenerationOutput.GenerationID)
    }
}