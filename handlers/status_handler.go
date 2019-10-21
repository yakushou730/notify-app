package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	oK = "ok"
)

// GetStatus returns 200 ok as response
func GetStatus(c *gin.Context) {
	c.String(http.StatusOK, oK)
	return
}
