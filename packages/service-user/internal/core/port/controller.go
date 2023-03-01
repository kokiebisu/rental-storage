package port

import customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"

type Controller interface {
	CreateUser(event interface{}) (interface{}, *customerror.CustomError)
	FindUserByEmail(event interface{}) (interface{}, *customerror.CustomError)
	FindUserById(event interface{}) (interface{}, *customerror.CustomError)
	RemoveUserById(event interface{}) (interface{}, *customerror.CustomError)
}
