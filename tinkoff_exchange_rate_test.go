package tinkoff_exchange_rate

import (
	"reflect"
	"testing"
)

func TestFetchCurrencyRatesHaveToReturnsErrNil(t *testing.T) {
	var tests = []struct {
		params map[string]string
	}{
		{
			params: map[string]string{},
		},
		{
			params: map[string]string{"from": "KZT", "to": "RUB"},
		},
	}

	for _, test := range tests {
		result, err := FetchCurrencyRates(test.params)

		if fromParamValue, ok := test.params["from"]; ok {
			for _, r := range result.Payload.Rates {
				if fromParamValue != r.FromCurrency.Name {
					t.Error("fromParamValue:", fromParamValue, " is not equal fetched rate ", r.FromCurrency.Name)
				}
			}
		}

		if toParamValue, ok := test.params["to"]; ok {
			for _, r := range result.Payload.Rates {
				if toParamValue != r.ToCurrency.Name {
					t.Error("toParamValue:", toParamValue, " is not equal fetched rate ", r.ToCurrency.Name)
				}
			}
		}

		if err != nil {
			t.Error("Function FetchCurrencyRates() have returned error: ", err)
		}
	}
}

func TestTableFilterRates(t *testing.T) {
	var tests = []struct {
		rates      []RateFromResponse
		filterFunc func(rate RateFromResponse) bool
		expected   []RateFromResponse
	}{
		{
			rates: []RateFromResponse{{
				Category: "any-category",
				FromCurrency: struct {
					Code    int    `json:"code"`
					Name    string `json:"name"`
					StrCode string `json:"strCode"`
				}{
					Code:    0,
					Name:    "USD",
					StrCode: "0",
				},
				ToCurrency: struct {
					Code    int    `json:"code"`
					Name    string `json:"name"`
					StrCode string `json:"strCode"`
				}{
					Code:    1,
					Name:    "AUD",
					StrCode: "1",
				},
				Buy:  10,
				Sell: 44,
			},
				{
					Category: "any-category",
					FromCurrency: struct {
						Code    int    `json:"code"`
						Name    string `json:"name"`
						StrCode string `json:"strCode"`
					}{
						Code:    0,
						Name:    "KZT",
						StrCode: "0",
					},
					ToCurrency: struct {
						Code    int    `json:"code"`
						Name    string `json:"name"`
						StrCode string `json:"strCode"`
					}{
						Code:    1,
						Name:    "AUD",
						StrCode: "1",
					},
					Buy:  10,
					Sell: 44,
				},
			},
			filterFunc: func(rate RateFromResponse) bool {
				return true
			},
			expected: []RateFromResponse{{
				Category: "any-category",
				FromCurrency: struct {
					Code    int    `json:"code"`
					Name    string `json:"name"`
					StrCode string `json:"strCode"`
				}{
					Code:    0,
					Name:    "USD",
					StrCode: "0",
				},
				ToCurrency: struct {
					Code    int    `json:"code"`
					Name    string `json:"name"`
					StrCode string `json:"strCode"`
				}{
					Code:    1,
					Name:    "AUD",
					StrCode: "1",
				},
				Buy:  10,
				Sell: 44,
			},
				{
					Category: "any-category",
					FromCurrency: struct {
						Code    int    `json:"code"`
						Name    string `json:"name"`
						StrCode string `json:"strCode"`
					}{
						Code:    0,
						Name:    "KZT",
						StrCode: "0",
					},
					ToCurrency: struct {
						Code    int    `json:"code"`
						Name    string `json:"name"`
						StrCode string `json:"strCode"`
					}{
						Code:    1,
						Name:    "AUD",
						StrCode: "1",
					},
					Buy:  10,
					Sell: 44,
				},
			},
		},
		{
			rates: []RateFromResponse{{
				Category: "any-category",
				FromCurrency: struct {
					Code    int    `json:"code"`
					Name    string `json:"name"`
					StrCode string `json:"strCode"`
				}{
					Code:    0,
					Name:    "USD",
					StrCode: "0",
				},
				ToCurrency: struct {
					Code    int    `json:"code"`
					Name    string `json:"name"`
					StrCode string `json:"strCode"`
				}{
					Code:    1,
					Name:    "AUD",
					StrCode: "1",
				},
				Buy:  10,
				Sell: 44,
			},
				{
					Category: "any-category",
					FromCurrency: struct {
						Code    int    `json:"code"`
						Name    string `json:"name"`
						StrCode string `json:"strCode"`
					}{
						Code:    0,
						Name:    "KZT",
						StrCode: "0",
					},
					ToCurrency: struct {
						Code    int    `json:"code"`
						Name    string `json:"name"`
						StrCode string `json:"strCode"`
					}{
						Code:    1,
						Name:    "AUD",
						StrCode: "1",
					},
					Buy:  10,
					Sell: 44,
				},
			},
			filterFunc: func(rate RateFromResponse) bool {
				return false
			},
			expected: []RateFromResponse{},
		},
	}

	for index, test := range tests {
		filtered := FilterRates(test.rates, test.filterFunc)

		if !reflect.DeepEqual(filtered, test.expected) {
			t.Error("Test #", index, " was failed")
		}
	}
}
