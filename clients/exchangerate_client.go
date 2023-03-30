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
	Success bool `json:"success"`
	Query   struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	} `json:"query"`
	Info struct {
		Rate float64 `json:"rate"`
	} `json:"info"`
	Historical bool    `json:"historical"`
	Date       string  `json:"date"`
	Result     float64 `json:"result"`
}

func getURI(client *ExchangeRatesClient, baseCurrency string, toCurrency string, amount string) string {
	uri := client.url + "/convert?from=" + baseCurrency + "&to=" + toCurrency
	if len(amount) != 0 {
		uri += "&amount=" + amount
	}
	return uri
}

// get the exchange rates from https://api.exchangerate.host
func (client *ExchangeRatesClient) getExchangeRateAmount(baseCurrency string, toCurrency string, amount string) (ExchangeRate, error) {
	var result ExchangeRate

	if client == nil {
		err := error(nil)
		return result, err
	}

	httpClient := http.Client{}

	uri := getURI(client, baseCurrency, toCurrency, amount)
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

	json.NewDecoder(resp.Body).Decode(&result)
	log.Debug(result)
	return result, nil
}

func convertExternalRateToRate(buyRate ExchangeRate, sellRate ExchangeRate, currency string) models.Rate {
	var result models.Rate
	result.Currency = currency
	result.Source = "exchangerate.host"
	result.UpdatedDate = time.Now().Format(time.RFC3339)
	result.BuyRate = -1
	result.SellRate = -1

	if buyRate != (ExchangeRate{}) {
		result.BuyRate = utils.SetFloatDecimalPoints(float64(buyRate.Info.Rate), 5)
	}

	if sellRate != (ExchangeRate{}) {
		result.SellRate = utils.SetFloatDecimalPoints(sellRate.Result, 5)
	}

	return result
}

func (client *ExchangeRatesClient) GetExchangeRateFromBase(base string, toCurrency string) (models.Rate, error) {
	log.Debug("converting from [" + base + "] to [ " + toCurrency + "]")

	// how many toCurrency you will get with 1 base
	buyRate, err := client.getExchangeRateAmount(base, toCurrency, "1")
	if err != nil {
		return convertExternalRateToRate(ExchangeRate{}, buyRate, toCurrency), err
	}

	// how many base you need for 1 toCurrency
	sellRate, err := client.getExchangeRateAmount(base, toCurrency, "1")
	if err != nil {
		return models.Rate{}, err
	}
	//buyRate.Result = 1000 / buyRate.Result

	result := convertExternalRateToRate(buyRate, sellRate, toCurrency)
	return result, nil
}

func (client *ExchangeRatesClient) GetExchangeRate(toCurrency string) (models.Rate, error) {
	return client.GetExchangeRateFromBase("DOP", toCurrency)
}
