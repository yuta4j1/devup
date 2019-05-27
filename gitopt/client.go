package gitopt

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func InitClient(accessToken string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}