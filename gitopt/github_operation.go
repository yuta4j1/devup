package gitopt

import (
	"fmt"
	"log"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	api "../api"
	. "../api/types"
)

func FetchAccessToken() (string, error) {
	// initialize response json data type
	authResponse := AuthorizationsResponse{
		Id: "",
		Url: "",
		Scopes: []string{""},
		Token: "",
		TokenLastEight: "",
		HashedToken: "",
		App: AppStruct{
			Url: "",
			Name: "",
			ClientId: "",
		},
		Note: "",
		NoteUrl: "",
		UpdatedAt: "",
		CreatedAt: "",
		Fingerprint: "",
	}
	// post
	out, err := api.Post("https://api.github.com/authorizations", AccessTokenParam{
		Scopes: []string{"repo"},
		Note: "get access token",
	}, authResponse)
	if err != nil {
		log.Fatal(err)
	}

	v, ok := out.(AuthorizationsResponse)
	if ok {
		fmt.Println(v)
	}

	return "", nil

}

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