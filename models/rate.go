package models

import "github.com/CryptocurrencyApp/CoinpocketServer/lib/rate"

func GetCoinRates() (result []rate.Rate, err error) {
	rates, err := rate.GetRates()
	result = rates.InfoList
	return
}
