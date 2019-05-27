package gitopt

import (
	"fmt"
	"log"
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