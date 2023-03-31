package controllers

import (
	"lmarrero/dop-exchange-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health
// @Accept */*
// @Produce json
// @Success 200 {object} models.Health
// @Router /health [get]
func (h HealthController) Status(c *gin.Context) {
	status := models.Health{Status: "UP"}
	c.IndentedJSON(http.StatusOK, status)
}
