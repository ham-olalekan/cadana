package services

import "os"

type ExchangeService interface {
	GetExchangeRate(currencyPair string) float64
}

var baseURL = os.Getenv("SERVICE_BASEURL")
