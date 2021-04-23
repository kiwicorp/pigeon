// cmd/lambda/main.go

package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/selftechio/pigeon/internal/content"
)

func handleRequest(ctx context.Context, params Params) (*Result, error) {
	var err error = nil
	err = validateParams(params)
	if err != nil {
		return nil, err
	}
	handler := content.NewGithubReleasesContentHandler(params.AccessToken, params.RepoOwner, params.RepoName)
	changed, err := handler.Handle(ctx)
	if err != nil {
		return nil, err
	}
	return newResult(changed), nil
}

func main() {
	lambda.Start(handleRequest)
}
