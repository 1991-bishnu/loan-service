package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Welcome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome")
}
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
