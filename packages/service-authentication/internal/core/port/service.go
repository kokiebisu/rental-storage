package port

type EncryptionService interface {
	SignIn(emailAddress string, password string) (string, error)
	SignUp(emailAddress string, firstName string, lastName string, password string) (string, error)
	Verify(authenticationToken string) (string, error)
}
