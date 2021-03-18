package src

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CurrencyExchangeHandler struct {
	E IExchangePriceModel
}

func NewCurrencyExchangeHandler(e IExchangePriceModel) *CurrencyExchangeHandler {
	return &CurrencyExchangeHandler{
		E: e,
	}
}

func (c *CurrencyExchangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fromQS := r.URL.Query()["from"]
	toQS := r.URL.Query()["to"]
	amountQS := r.URL.Query()["amount"]

	w.Header().Add("content-type", "application/json")
	if len(fromQS) == 0 || len(toQS) == 0 || len(amountQS) == 0{
		w.WriteHeader(400)
		errorRes := struct{
			Error string
		}{
			Error: "Need from/to/amount",
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}
	from:=fromQS[0]
	to:=toQS[0]
	amount,_ := strconv.Atoi(amountQS[0])

	exchageAmount := c.Exchange(from, to, amount)
	w.WriteHeader(200)
	resBody := struct {
		Amount int
	}{
		Amount: exchageAmount,
	}
	json.NewEncoder(w).Encode(resBody)
}

func (c *CurrencyExchangeHandler) Exchange(from string, to string, amount int) int {
	if amount == 0 {
		return 0
	}

	ch := make(chan ExchangeRateResult)
	go c.E.GetExchangeRate(from, to, ch)
	result := <-ch
	if result.IsExists {
		return amount * result.ExchangeRate
	}

	return 0
}
