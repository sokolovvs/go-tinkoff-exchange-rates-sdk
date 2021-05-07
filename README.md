# Golang SDK for Tinkoff Bank currency exchange rates
[![Build Status](https://travis-ci.com/sokolovvs/go-tinkoff-exchange-rates-sdk.svg?branch=master)](https://travis-ci.com/sokolovvs/go-tinkoff-exchange-rates-sdk)

### Usage example

```go
func UpdateTinkoffRates() {
	defaultFilterFunc := func(rate RateFromResponse) bool {
		if  rate.Category == "C2CTransfers" {
			return true
		}

		return false
	}

	UpdateTinkoffRatesByParams(map[string]string{"from": "USD", "to": "RUB"}, defaultFilterFunc)
	UpdateTinkoffRatesByParams(map[string]string{"from": "EUR", "to": "RUB"}, defaultFilterFunc)
	UpdateTinkoffRatesByParams(map[string]string{"from": "KZT", "to": "RUB"}, defaultFilterFunc)
	UpdateTinkoffRatesByParams(map[string]string{"from": "CAD", "to": "RUB"}, defaultFilterFunc)
	UpdateTinkoffRatesByParams(map[string]string{"from": "AUD", "to": "RUB"}, defaultFilterFunc)
}

func UpdateTinkoffRatesByParams(params map[string]string, filterFunc func(response RateFromResponse) bool) {
	response, err := FetchCurrencyRates(params)

	if err != nil {
		log.Error(err)
		return
	}

	rates := FilterRates(response.Payload.Rates, filterFunc)
}
```