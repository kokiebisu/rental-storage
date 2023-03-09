package test

import (
	"fmt"
	"testing"

	"github.com/kokiebisu/rental-storage/service-booking/test/data"

	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-booking/internal/helper"

	"github.com/stretchr/testify/assert"
)

type StringifyResult struct {
	input    interface{}
	expected string
}

var stringifyResults = []StringifyResult{
	{
		adapter.CreateBookingResponsePayload{
			UId: data.MockBooking.UId,
		},
		fmt.Sprintf(`{"uid":"%s"}`, data.MockBooking.UId),
	},
}

func TestStringify(t *testing.T) {
	for _, test := range stringifyResults {
		result, _ := helper.Stringify(test.input)
		assert.Equal(t, test.expected, result)
	}
}
