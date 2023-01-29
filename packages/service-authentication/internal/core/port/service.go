package port

import errors "github.com/kokiebisu/rental-storage/service-authentication/internal/error"

type EncryptionService interface {
	SignIn(emailAddress string, password string) (string, *errors.CustomError)
	SignUp(emailAddress string, firstName string, lastName string, password string) (string, *errors.CustomError)
	Verify(authenticationToken string) (string, *errors.CustomError)
}
