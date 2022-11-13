package port

type SignInArgument struct {
	EmailAddress      string `json:"emailAddress"`
	Password   string `json:"password"`
}