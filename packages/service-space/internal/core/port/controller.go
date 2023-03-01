package port

import (
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type Controller interface {
	FindSpaceById(event interface{}) (interface{}, *customerror.CustomError)
	FindSpaces(event interface{}) (interface{}, *customerror.CustomError)
	AddSpace(event interface{}) (interface{}, *customerror.CustomError)
	DeleteSpaceById(event interface{}) (interface{}, *customerror.CustomError)
}
