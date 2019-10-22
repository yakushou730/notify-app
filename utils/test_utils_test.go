package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMockedContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/test", nil)
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockedContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "8080", c.Request.URL.Port())
	assert.EqualValues(t, "/test", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
}
