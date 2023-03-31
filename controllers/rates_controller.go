package controllers

import (
	"lmarrero/dop-exchange-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RatesController struct{}

var ratesService = new(service.RatesService)

// GetRateFromDOP godoc
// @Summary Get the exchange rates from DOP to the target currency
// @Schemes
// @Description Get the exchange rates from DOP to the target currency
// @Tags Rates
// @Produce json
// @Success 200 {object} []models.Rate
// @Router /v1/rates/:targetCurrency [get]
func (r RatesController) GetRateFromDOP(c *gin.Context) {
	rate, err := ratesService.GetFromDOP(c.Param("target"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to get rate")
	}
	c.IndentedJSON(http.StatusOK, rate)
	c.Abort()
	return
}

// GetAllDOPRates godoc
// @Summary Get rates from BASE to TARGET
// @Schemes
// @Description Get rates from BASE to TARGET
// @Tags Rates
// @Produce json
// @Success 200 {object} []models.Rate
// @Router /v1/rates/custom/:baseCurrency/:targetCurrency [get]
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

// GetAllDOPRates godoc
// @Summary Get rates from DOP to all supported currencies
// @Schemes
// @Description Get rates from DOP to all supported currencies
// @Tags Rates
// @Produce json
// @Success 200 {object} []models.Rate
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
