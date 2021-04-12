// cmd/lambda/main.go

package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/selftechio/pigeon/internal/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GithubReleasesParams struct {
	AccessToken        string `json:"access_token"`
	RepoOwner          string `json:"repo_owner"`
	RepoName           string `json:"repo_name"`
	PreviousTotalCount int32  `json:"prev_total_count"`
}

func handleRequest(ctx context.Context, params GithubReleasesParams) (string, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: params.AccessToken},
	)
	httpClient := oauth2.NewClient(ctx, src)
	client := githubv4.NewClient(httpClient)

	totalCount := github.RepositoryReleasesTotalCountQuery{}
	variables := map[string]interface{}{
		"owner": githubv4.String(params.RepoOwner),
		"name":  githubv4.String(params.RepoName),
	}
	err := client.Query(ctx, &totalCount, variables)
	if err != nil {
		return "", err
	}

	releases := github.RepositoryReleasesQuery{}
	variables = map[string]interface{}{
		"owner": githubv4.String("tokio-rs"),
		"name":  githubv4.String("tokio"),
		"first": githubv4.Int(int(totalCount.Repository.Releases.TotalCount) - int(params.PreviousTotalCount)),
	}
	err = client.Query(ctx, &releases, variables)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Query response: %#v", releases.Repository.Releases.Edges), nil
}

func main() {
	lambda.Start(handleRequest)
}
