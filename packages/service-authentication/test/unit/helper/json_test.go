package test

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/helper"
	"github.com/stretchr/testify/assert"
)

type StringifyResult struct {
	input    interface{}
	expected string
}

var stringifyResults = []StringifyResult{
	{
		payload{
			UId: "random",
		},
		`{"uid":"random"}`,
	},
	{
		"",
		`""`,
	},
}

type payload struct {
	UId string `json:"uid"`
}

// for a valid return value.
func TestStringify(t *testing.T) {
	for _, test := range stringifyResults {
		result, _ := helper.Stringify(test.input)
		assert.Equal(t, test.expected, result)
	}
}
