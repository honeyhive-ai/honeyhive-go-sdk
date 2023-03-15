<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/honeyhive-ai/honeyhive-go-sdk"
    "github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
    "github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/operations"
)

func main() {
    s := honeyhive.New(
        WithSecurity(        shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    req := shared.Dataset{
        Description: "unde",
        File: "deserunt",
        Name: "porro",
        Purpose: "nulla",
        Task: "id",
    }

    ctx := context.Background()
    res, err := s.Dataset.CreateDataset(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDatasetResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->