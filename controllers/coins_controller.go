package controllers

import (
	"github.com/gin-gonic/gin"
	"lmarrero/dop-exchange-api/models"
	"net/http"
)

type CoinsController struct{}

func (cc CoinsController) GetSupportedCoins(c *gin.Context) {
	coins := models.SupportedCurrencies()
	c.IndentedJSON(http.StatusOK, coins)
}
