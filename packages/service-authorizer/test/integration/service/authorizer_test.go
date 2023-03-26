package integration

import (
	"testing"

	"github.com/kokiebisu/rental-storage/service-authorizer/test/data"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize_Success(t *testing.T) {
	claim, err := data.AuthorizerService.Authorize(data.MockToken)
	assert.Nil(t, err, "should not throw error")
	assert.NotNil(t, claim, "should not be nil")
}
