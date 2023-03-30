package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"lmarrero/dop-exchange-api/clients"
	"lmarrero/dop-exchange-api/models"
)

type RatesController struct{}

func getFromExternal(currency string) models.Rate {
	client := clients.Init("https://api.exchangerate.host")
	// get sell rate from external
	rate, err := client.GetExchangeRate(currency)

	if err != nil {
		return models.Rate{}
	}

	return rate
}

func getRates(currency string) ([]models.Rate, error) {
	var rates []models.Rate

	// get from external api
	externalRate := getFromExternal(currency)
	if (externalRate != models.Rate{}) {
		rates = append(rates, externalRate)
	}

	return rates, nil
}

func (r RatesController) GetRate(c *gin.Context) {
	rate, err := getRates(c.Param("coin"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to get rate")
	}
	c.IndentedJSON(http.StatusOK, rate)
	c.Abort()
	return
}

func (r RatesController) GetAllRates(c *gin.Context) {
	var result []models.Rate

	for _, currency := range models.SupportedCurrencies() {
		rate, err := getRates(currency)
		if err != nil {
			fmt.Println("Failed to get " + currency + " rate")
		} else {
			result = append(result, rate...)
		}
	}

	eurRates, err := getRates("EUR")
	if err != nil {
		fmt.Println("Failed to get EUR rate")
	}

	result = append(result, eurRates...)

	c.IndentedJSON(http.StatusOK, result)
	c.Abort()
	return
}
