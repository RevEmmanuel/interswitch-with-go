package authService

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorization(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Test GetToken function
	accessToken := GetToken()
	assert.NotNil(t, accessToken, "The access token should not be nil")
	assert.NotEmpty(t, accessToken, "The access token should not be empty")
	t.Logf("Access Token is: %s", accessToken)
}
