package types

type AccessTokenParam struct {
	Scopes []string `json:scopes`
	Note string    `json:note`
}

type AppStruct struct {
	Url string `json:url`
	Name string `json:name`
	ClientId string `json:client_id`
}

type AuthorizationsResponse struct {
	Id string `json:id`
	Url string `json:url`
	Scopes []string `json:scopes`
	Token string `json:token`
	TokenLastEight string `json:token_last_eight`
	HashedToken string `json:hashed_token`
	App AppStruct `json:app`
	Note string `json:note`
	NoteUrl string `json:note_url`
	UpdatedAt string `json:updated_at`
	CreatedAt string `json:created_at`
	Fingerprint string `json:fingerprint`

}