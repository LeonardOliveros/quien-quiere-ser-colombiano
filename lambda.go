package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

// seedActionPayload is the direct-invoke payload used by the deploy-time
// seeder (a CDK custom resource calls lambda:Invoke with it):
//
//	{"quizapp_action": "seed"}
type seedActionPayload struct {
	Action string `json:"quizapp_action"`
}

// runLambda serves API Gateway HTTP API (payload v2) events through the Gin
// router, and handles the special seed invoke used at deploy time. CloudFront
// forwards /api/* unchanged and the HTTP API routes ANY /{proxy+}, so the
// paths reaching Gin are identical to the local server's — no rewriting.
func runLambda(r *gin.Engine) {
	adapter := ginadapter.NewV2(r)
	lambda.Start(func(ctx context.Context, raw json.RawMessage) (any, error) {
		var action seedActionPayload
		if err := json.Unmarshal(raw, &action); err == nil && action.Action == "seed" {
			log.Println("Seed invoke: syncing question bank")
			if err := syncQuestionBank(); err != nil {
				return nil, fmt.Errorf("seed: %w", err)
			}
			return map[string]string{"status": "seeded"}, nil
		}

		var req events.APIGatewayV2HTTPRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, fmt.Errorf("unsupported event payload: %w", err)
		}
		return adapter.ProxyWithContext(ctx, req)
	})
}
