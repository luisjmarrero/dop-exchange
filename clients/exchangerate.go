package clients

import (
	"encoding/json"
	"lmarrero/dop-exchange-api/models"
	"lmarrero/dop-exchange-api/utils"
	"net/http"
	"time"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("client")

type ExchangeRatesClient struct {
	url string
}

func Init(url string) *ExchangeRatesClient {
	client := &ExchangeRatesClient{
		url: url,
	}
	return client
}

type ExchangeRate struct {
	Success bool    `json:"success"`
	Result  float64 `json:"result"`
	Date    string  `json:"date"`
}

func getURI(client *ExchangeRatesClient, baseCurrency string, toCurrency string) string {
	uri := client.url + "/convert?from=" + baseCurrency + "&to=" + toCurrency
	return uri
}

func externalRateToRate(buyRate ExchangeRate, sellRate ExchangeRate, currency string) models.Rate {
	var result models.Rate
	result.Currency = currency
	result.Source = "exchangerate.host"
	result.UpdatedDate = time.Now().Format(time.RFC3339)
	result.BuyRate = -1
	result.SellRate = -1

	if buyRate != (ExchangeRate{}) {
		result.BuyRate = utils.SetFloatDecimalPoints(buyRate.Result, 5)
	}

	if sellRate != (ExchangeRate{}) {
		result.SellRate = utils.SetFloatDecimalPoints(sellRate.Result, 5)
	}

	return result
}

// get the exchange rates from https://api.exchangerate.host
func (client *ExchangeRatesClient) getExchangeRateAmount(baseCurrency string, toCurrency string) (ExchangeRate, error) {
	var result ExchangeRate

	if client == nil {
		err := error(nil)
		return result, err
	}

	httpClient := http.Client{}

	uri := getURI(client, baseCurrency, toCurrency)
	log.Debug("URI= " + uri)
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		// fmt.Println(err)
		log.Error(err)
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return result, err
	}

	log.Debug(resp.Body)

	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (client *ExchangeRatesClient) GetExchangeRate(toCurrency string) (models.Rate, error) {
	log.Debug("toCurrency: " + toCurrency)
	sellRate, err := client.getExchangeRateAmount("DOP", toCurrency)
	if err != nil {
		return models.Rate{}, err
	}

	buyRate, err := client.getExchangeRateAmount(toCurrency, "DOP")
	// TODO: get reverse exchange.. ?
	if err != nil {
		return externalRateToRate(ExchangeRate{}, sellRate, toCurrency), err
	}

	result := externalRateToRate(buyRate, sellRate, toCurrency)
	return result, nil
}
