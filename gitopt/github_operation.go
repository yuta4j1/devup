package gitopt

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func InitClient(accessToken string) (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), ctx
}

func GithubCreateRepository(client *github.Client, ctx context.Context, projName string) (*github.Repository, *github.Response, error) {
	newRepo, response, err := client.Repositories.Create(ctx, "", &github.Repository{
		Name: &projName,
	})
	return newRepo, response, err
}