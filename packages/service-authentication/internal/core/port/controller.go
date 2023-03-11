package port

import customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"

type Controller interface {
	SignIn(req interface{}) (interface{}, *customerror.CustomError)
	SignUp(req interface{}) (interface{}, *customerror.CustomError)
	Verify(req interface{}) (interface{}, *customerror.CustomError)
}
