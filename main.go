package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Rate struct {
	Success bool    `json:"success"`
	Result  float64 `json:"result"`
	Date    string  `json:"date"`
}

type ExchangeRate struct {
	Rate        float64 `json:"rate"`
	UpdatedDate string  `json:"updated_date"`
	Source      string  `json:"source"`
}

type ExchangeRates struct {
	BaseCurrency string                    `json:"base"`
	Rates        map[string][]ExchangeRate `json:"rates"`
}

// get the exchange rates from https://api.exchangerate.host
func getExchangeRatesFromExternalAPI(baseCurrency string) (*Rate, error) {
	client := http.Client{}

	request, err := http.NewRequest("GET", "https://api.exchangerate.host/convert?from="+baseCurrency+"&to=DOP", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)
	var result Rate
	json.NewDecoder(resp.Body).Decode(&result)
	return &result, nil
}

func setTwoDecimalPoints(x float64) float64 {
	return float64(int(x*100)) / 100
}

func rateToExchangeRate(rate *Rate) *ExchangeRate {
	if rate == nil {
		return nil
	}

	strDate, _ := time.Parse(time.DateOnly, rate.Date)

	result := ExchangeRate{
		Rate:        setTwoDecimalPoints(rate.Result),
		UpdatedDate: strDate.Format(time.RFC3339),
		Source:      "exchangerate.host",
	}

	return &result
}

func getRateExternal(destination string) (*ExchangeRate, error) {
	rate, err := getExchangeRatesFromExternalAPI(destination)
	if err != nil {
		return nil, err
	}
	exchangeRate := rateToExchangeRate(rate)
	return exchangeRate, nil
}

func getRates(destination string) ([]ExchangeRate, error) {
	var rates []ExchangeRate

	// get from external
	externalRate, _ := getRateExternal(destination)
	rates = append(rates, *externalRate)

	return rates, nil
}

func getUSDRates(c *gin.Context) {
	rate, err := getRates("USD")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to get USD rate")
	}
	c.IndentedJSON(http.StatusOK, rate)
}

func getEURRates(c *gin.Context) {
	rate, err := getRates("EUR")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to get EUR rate")
	}
	c.IndentedJSON(http.StatusOK, rate)
}

func getAllRates(c *gin.Context) {
	var result ExchangeRates
	result.BaseCurrency = "DOP"

	usdRates, err := getRates("USD")
	if err != nil {
		fmt.Println("Failed to get USD rate")
	}

	eurRates, err := getRates("EUR")
	if err != nil {
		fmt.Println("Failed to get EUR rate")
	}

	result.Rates = map[string][]ExchangeRate{
		"USD": usdRates,
		"EUR": eurRates,
	}

	c.IndentedJSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()
	router.GET("/rates/USD", getUSDRates)
	router.GET("/rates/EUR", getEURRates)
	router.GET("/rates", getAllRates)

	router.Run(":8080")
}
