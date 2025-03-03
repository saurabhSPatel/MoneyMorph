package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const API_URL = "https://v6.exchangerate-api.com/v6/<API_KEY>/latest/"

type ExchangeRates struct {
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func getExchangeRate(baseCurr string) (ExchangeRates, error) {
	var rates ExchangeRates
	resp, err := http.Get(API_URL + baseCurr)
	if err != nil {
		return rates, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		return rates, err
	}
	return rates, nil
}
func ConvertCurrency(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	base := query.Get("base")
	target := query.Get("target")
	amount := query.Get("amount")

	if base == "" || target == "" || amount == "" {
		http.Error(w, "Missing query parameters: base, target, amount", http.StatusBadRequest)
		return
	}

	rates, err := getExchangeRate(base)
	if err != nil {
		http.Error(w, "Failed to fetch exchange rates", http.StatusInternalServerError)
		return
	}

	exchangeRate, exists := rates.ConversionRates[target]
	if !exists {
		http.Error(w, "Invalid target currency", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"base":    base,
		"target":  target,
		"rate":    exchangeRate,
		"message": fmt.Sprintf("1 %s = %.2f %s", base, exchangeRate, target),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
