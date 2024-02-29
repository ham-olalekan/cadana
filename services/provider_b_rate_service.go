package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ham-olalekan/cadana/secrets"
	"io/ioutil"
	"net/http"
)

type ServiceB struct {
	SecretData secrets.SecretData
}

func (s ServiceB) GetExchangeRate(currencyPair string) (rate float64, err error) {
	apiKey := s.SecretData.ApikeyProviderA
	payload := map[string]string{
		"currencyPair": currencyPair,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding payload:", err)
		return
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", baseURL, body)
	if err != nil {
		fmt.Println("Error creating http request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(respBody))
	return
}
