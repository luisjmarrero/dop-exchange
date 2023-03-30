package service

import (
	"fmt"
	"lmarrero/dop-exchange-api/clients"
	"lmarrero/dop-exchange-api/models"
)

type RatesService struct{}

func getExternalRateFromBase(base string, to string) models.Rate {
	client := clients.Init("https://api.exchangerate.host")
	// get sell rate from external
	rate, err := client.GetExchangeRateFromBase(base, to)

	if err != nil {
		return models.Rate{}
	}

	return rate
}

func (s *RatesService) GetRatesFromBase(base string, to string) ([]models.Rate, error) {
	var rates []models.Rate

	// get from external api
	externalRate := getExternalRateFromBase(base, to)
	if (externalRate != models.Rate{}) {
		rates = append(rates, externalRate)
	}

	// eventually add more sources

	return rates, nil
}

func (s *RatesService) GetFromDOP(target string) ([]models.Rate, error) {
	return s.GetRatesFromBase("DOP", target)
}

func (s *RatesService) GetAllDOPRates() ([]models.Rate, error) {
	var result []models.Rate
	for _, currency := range models.SupportedCurrencies() {
		rate, err := s.GetFromDOP(currency)
		if err != nil {
			fmt.Println("Failed to get " + currency + " rate")
		} else {
			result = append(result, rate...)
		}
	}
	return result, nil
}
