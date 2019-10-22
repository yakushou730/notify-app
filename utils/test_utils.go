package utils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// GetMockedContext is used for gin context mock
func GetMockedContext(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return c
}
