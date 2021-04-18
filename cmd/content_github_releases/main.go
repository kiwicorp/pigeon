// cmd/lambda/main.go

package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/selftechio/pigeon/internal/content/github_releases"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func handleRequest(ctx context.Context, params Params) (*Result, error) {
	var err error

	err = validateParams(params)
	if err != nil {
		return nil, err
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: params.AccessToken},
	)
	httpClient := oauth2.NewClient(ctx, src)
	client := githubv4.NewClient(httpClient)

	totalCount := github_releases.RepositoryReleasesTotalCountQuery{}
	variables := map[string]interface{}{
		"owner": githubv4.String(params.RepoOwner),
		"name":  githubv4.String(params.RepoName),
	}
	err = client.Query(ctx, &totalCount, variables)
	if err != nil {
		return nil, err
	}

	prevTotalCount := 0

	releases := github_releases.RepositoryReleasesQuery{}
	variables = map[string]interface{}{
		"owner": githubv4.String("tokio-rs"),
		"name":  githubv4.String("tokio"),
		"first": githubv4.Int(int(totalCount.Repository.Releases.TotalCount) - prevTotalCount),
	}
	err = client.Query(ctx, &releases, variables)
	if err != nil {
		return nil, err
	}

	// return fmt.Sprintf("Query response: %#v", releases.Repository.Releases.Edges), nil
	return newResult(0), nil // fixme 17/04/2021: replace with proper return value
}

func main() {
	lambda.Start(handleRequest)
}
