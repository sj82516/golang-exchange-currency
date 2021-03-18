package src

type IExchangePriceModel interface {
	GetExchangeRate(string, string, chan<- ExchangeRateResult)
}

type ExchangePriceModel struct {
}

func (m *ExchangePriceModel) GetExchangeRate(from string, to string, ch chan<- ExchangeRateResult) {
	// TODO: replace by real db query
	if from == "US" && to == "TW" {
		ch <- ExchangeRateResult{
			IsExists:     true,
			ExchangeRate: 30,
		}
		return
	}
	ch <- ExchangeRateResult{
		IsExists:     false,
		ExchangeRate: 0,
	}
}

type ExchangeRateResult struct {
	IsExists     bool
	ExchangeRate int
}
