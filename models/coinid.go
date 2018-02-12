package models

import "github.com/CryptocurrencyApp/CoinpocketServer/lib/rate"

type CoinId struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func GetCoinIds() (result []CoinId, err error) {

	rates, err := rate.GetRates()
	for _, r := range rates.InfoList {
		result = append(result, CoinId{
			Id:     r.ID,
			Name:   r.Name,
			Symbol: r.Symbol,
		})
	}

	return
}
