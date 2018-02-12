package rate

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

const priceFilePath = "./rateLog/newst.json"

type Rates struct {
	GetAt    string
	InfoList []Rate
}

type Rate struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	PriceUsd         float64 `json:"price_usd"`
	PriceJpy         float64 `json:"price_jpy"`
	PriceBtc         float64 `json:"price_btc"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
}

func GetJpyPrice(coinId string, rates Rates) (price string) {
	for _, rate := range rates.InfoList {
		if rate.ID == coinId {
			price = fmt.Sprint(rate.PriceJpy)
		}
	}

	return price
}

func GetRates() (rates Rates, err error) {
	file, err := os.OpenFile(priceFilePath, os.O_RDONLY, 700)
	if err != nil {
		return
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(raw, &rates)
	if err != nil {
		return
	}

	return
}
