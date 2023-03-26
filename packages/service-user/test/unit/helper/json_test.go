package unit

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"

	"github.com/kokiebisu/rental-storage/service-user/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/helper"
	"github.com/stretchr/testify/assert"
)

type StringifyResult struct {
	input    interface{}
	expected string
}

var (
	UId          = faker.UUIDDigit()
	EmailAddress = faker.Email()
	FirstName    = faker.FirstName()
	LastName     = faker.LastName()
	Password     = faker.Password()
	TimeString   = faker.TimeString()
)

var stringifyResults = []StringifyResult{
	{
		adapter.CreateUserResponsePayload{
			UId: UId,
		},
		fmt.Sprintf(`{"uid":"%s"}`, UId),
	},
	{
		adapter.FindUserByEmailResponsePayload{
			User: user.DTO{
				FirstName:    FirstName,
				LastName:     LastName,
				EmailAddress: EmailAddress,
				Password:     Password,
				Items:        []item.DTO{},
				CreatedAt:    TimeString,
				UpdatedAt:    TimeString,
			},
		},
		fmt.Sprintf(`{"user":{"uid":"%s","firstName":"%s","lastName":"%s","emailAddress":"%s","password":"%s","items":[],"createdAt":"%s","updatedAt":"%s"}}`, "", FirstName, LastName, EmailAddress, Password, TimeString, TimeString),
	},
}

func TestStringify(t *testing.T) {
	for _, test := range stringifyResults {
		result, _ := helper.Stringify(test.input)
		assert.Equal(t, test.expected, result)
	}
}
