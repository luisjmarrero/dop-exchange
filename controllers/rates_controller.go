package controllers

import (
	"lmarrero/dop-exchange-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RatesController struct{}

var ratesService = new(service.RatesService)

func (r RatesController) GetRateFromDOP(c *gin.Context) {
	rate, err := ratesService.GetFromDOP(c.Param("target"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to get rate")
	}
	c.IndentedJSON(http.StatusOK, rate)
	c.Abort()
	return
}

func (r RatesController) GetRateFromBase(c *gin.Context) {
	base := c.Param("base")
	to := c.Param("target")
	rate, err := ratesService.GetRatesFromBase(base, to)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to get rate")
	}
	c.IndentedJSON(http.StatusOK, rate)
	c.Abort()
	return
}

// @BasePath /api/v1

// PingExample godoc
// @Summary et Rates from DOP to all supported currencies
// @Schemes
// @Description Get Rates from DOP to all supported currencies
// @Tags rates
// @Produce json
// @Success 200
// @Router /v1/rates/ [get]
func (r RatesController) GetAllDOPRates(c *gin.Context) {
	rates, err := ratesService.GetAllDOPRates()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to get rate")
	}

	c.IndentedJSON(http.StatusOK, rates)
	c.Abort()
	return
}
