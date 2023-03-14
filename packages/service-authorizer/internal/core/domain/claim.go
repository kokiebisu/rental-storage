package domain

type Claim struct {
	UId string `json:"uid"`
	Exp int    `json:"exp"`
}
