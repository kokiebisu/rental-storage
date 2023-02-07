package test

import (
	"fmt"
	"testing"

	"github.com/kokiebisu/rental-storage/service-authentication/test/data"
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/controller"
	"github.com/kokiebisu/rental-storage/service-booking/internal/helper"

	"github.com/stretchr/testify/assert"
)

type StringifyResult struct {
	input    interface{}
	expected string
}

var stringifyResults = []StringifyResult{
	{
		controller.CreateBookingResponsePayload{
			UId: data.MockUId,
		},
		fmt.Sprintf(`{"uid":"%s"}`, data.MockUId),
	},
}

func TestStringify(t *testing.T) {
	for _, test := range stringifyResults {
		result, _ := helper.Stringify(test.input)
		assert.Equal(t, test.expected, result)
	}
}
