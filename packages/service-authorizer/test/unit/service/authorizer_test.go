package unit

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kokiebisu/rental-storage/service-authorizer/test/data"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := fmt.Sprintf(`{"uid":"%s","exp":%d}`, data.MockUId, data.MockExpiresAt)
		_, err := w.Write([]byte(response))
		if err != nil {
			log.Fatal("Unable to write to body")
		}
	}))
	defer server.Close()
	os.Setenv("SERVICE_API_ENDPOINT", server.URL)

	claim, err := data.AuthorizerService.Authorize(data.MockToken)
	assert.Nil(t, err)
	assert.Equal(t, data.MockUId, claim.UId)
}
