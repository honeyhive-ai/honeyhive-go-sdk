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
    generationID := "63ffe425878a42b6aefead62"
	feedbackJSON := map[string]interface{}{
		"feedback": "This is a test",
		"accepted": true,
	}
    
    req := operations.CreateFeedbackRequest{
        Request: shared.Feedback{
            Task: &taskName,
			GenerationID: &generationID,
			FeedbackJSON: feedbackJSON,
        },
    }

    ctx := context.Background()
    res, err := s.Feedback.CreateFeedback(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Feedback != nil {
        // handle response
		fmt.Println(*res.Feedback.GenerationID)
    }
}