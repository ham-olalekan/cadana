package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
	"time"
)

type ExchangeRateResponse struct {
	CurrencyPair string  `json:"currencyPair"`
	Rate         float64 `json:"rate"`
}

type ExchangeRateRequest struct {
	CurrencyPair string `json:"currencyPair"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/exchange-rate", exchangeRateHandler).Methods("POST")

	// Start the server with Gorilla Mux as the handler
	http.ListenAndServe(":8081", r)
}

func exchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	var request ExchangeRateResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	var rate float64
	var once sync.Once

	wg.Add(1)
	go func() {
		defer wg.Done()
		//call service A
		tempRate := GetExchangeRateFromServiceA(request.CurrencyPair)
		once.Do(func() {
			rate = tempRate
		})
	}()
	go func() {
		defer wg.Done()
		//call service B
		tempRate := GetExchangeRateFromServiceB(request.CurrencyPair)
		once.Do(func() {
			rate = tempRate
		})
	}()
	wg.Wait()

	response := ExchangeRateResponse{CurrencyPair: request.CurrencyPair, Rate: rate}
	json.NewEncoder(w).Encode(response)
}

func GetExchangeRateFromServiceA(currencyPair string) float64 {
	time.Sleep(4 * time.Second)
	return 0.91
}

func GetExchangeRateFromServiceB(currencyPair string) float64 {
	time.Sleep(2 * time.Second)
	return 0.92
}
