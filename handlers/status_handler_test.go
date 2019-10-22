package handlers

import (
	"net/http"
	"net/http/httptest"
	"notify-app/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "ok", oK)
}

func TestGetStatus(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/status", nil)
	c := utils.GetMockedContext(request, response)

	GetStatus(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "ok", response.Body.String())
}
