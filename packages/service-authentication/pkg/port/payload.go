package port

type UserCreationPayload struct {
	EmailAddress      string `json:"emailAddress"`
	Password   string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type AuthorizationTokenPayload struct {
	AuthorizationToken string `json:"authorizationToken"`
}

type GenerateJWTTokenPayload struct {
    UId string `json:"uid"`
}
