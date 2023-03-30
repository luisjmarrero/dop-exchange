package controllers

import (
	"lmarrero/dop-exchange-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	status := models.Health{Status: "UP"}
	c.IndentedJSON(http.StatusOK, status)
}
