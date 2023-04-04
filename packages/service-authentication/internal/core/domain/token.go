package domain

type Token string

type Tokens struct {
	AccessToken  Token
	RefreshToken Token
}
