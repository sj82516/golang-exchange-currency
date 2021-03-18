package main

import (
	"currency-exchange/src"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	exchangePriceModel := &src.ExchangePriceModel{}
	currencyExchangeHandler := &src.CurrencyExchangeHandler{
		E: exchangePriceModel,
	}
	mux.HandleFunc("/exchange-currency", currencyExchangeHandler.ServeHTTP)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
