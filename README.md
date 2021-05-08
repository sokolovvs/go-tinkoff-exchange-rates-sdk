# Golang SDK for Tinkoff Bank currency exchange rates

[![Build Status](https://travis-ci.com/sokolovvs/go-tinkoff-exchange-rates-sdk.svg?branch=master)](https://travis-ci.com/sokolovvs/go-tinkoff-exchange-rates-sdk)
[![Go](https://github.com/sokolovvs/go-tinkoff-exchange-rates-sdk/actions/workflows/go.yml/badge.svg)](https://github.com/sokolovvs/go-tinkoff-exchange-rates-sdk/actions/workflows/go.yml)

### Package installation

```shell
go get github.com/sokolovvs/go-tinkoff-exchange-rates-sdk
```

### Usage example

```go
package main

import (
	"fmt"
	tinkoff_exchange_rate "github.com/sokolovvs/go-tinkoff-exchange-rates-sdk"
)

func main() {
	UpdateTinkoffRates()
}

func UpdateTinkoffRates() {
	defaultFilterFunc := func(rate tinkoff_exchange_rate.RateFromResponse) bool {
		if rate.Category == "C2CTransfers" {
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

func UpdateTinkoffRatesByParams(params map[string]string, filterFunc func(response tinkoff_exchange_rate.RateFromResponse) bool) {
	response, err := tinkoff_exchange_rate.FetchCurrencyRates(params)

	if err != nil {
		panic(err)
		return
	}

	rates := tinkoff_exchange_rate.FilterRates(response.Payload.Rates, filterFunc)

	fmt.Println(rates)
}
```
