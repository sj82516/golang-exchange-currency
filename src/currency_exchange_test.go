package src

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("currency exchange", func() {
	c := &CurrencyExchangeHandler{}
	var (
		mockCtrl *gomock.Controller
		e        *MockIExchangePriceModel
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		e = NewMockIExchangePriceModel(mockCtrl)
		c = NewCurrencyExchangeHandler(e)
	})

	It("should get 0 if amount is 0", func(done Done) {
		ch := make(chan ExchangeRateResult)
		e.EXPECT().GetExchangeRate("US", "TW", ch).Do(func() {
			ch <- ExchangeRateResult{
				IsExists:     true,
				ExchangeRate: 40,
			}
		})
		Expect(c.Exchange("US", "TW", 0)).To(Equal(0))
		close(done)
	})

	It("get exchange price table and calc exchange price", func(done Done) {
		e.EXPECT().GetExchangeRate("US", "TW", gomock.Any()).Do(func(from string, to string, ch chan<- ExchangeRateResult) {
			ch <- ExchangeRateResult{
				IsExists:     true,
				ExchangeRate: 30,
			}
			close(ch)
		})
		Expect(c.Exchange("US", "TW", 30)).To(Equal(900))
		close(done)
	})

	It("return 0 if the exchange currency mapping is not support", func(done Done) {
		e.EXPECT().GetExchangeRate(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(from string, to string, ch chan<- ExchangeRateResult) {
			ch <- ExchangeRateResult{
				IsExists:     false,
				ExchangeRate: 0,
			}
			close(ch)
		})
		Expect(c.Exchange("Non", "TW", 30)).To(Equal(0))
		close(done)
	})

	It("test ServeHttp integration", func(done Done) {
		e.EXPECT().GetExchangeRate(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(from string, to string, ch chan<- ExchangeRateResult) {
			ch <- ExchangeRateResult{
				IsExists:     false,
				ExchangeRate: 0,
			}
			close(ch)
		})

		req,_ := http.NewRequest("GET", "/exchange-currency", nil)
		query := req.URL.Query()
		query.Add("from", "US")
		query.Add("to", "TW")
		query.Add("amount", "10")
		req.URL.RawQuery = query.Encode()

		rr := httptest.NewRecorder()
		e := &ExchangePriceModel{}
		c := &CurrencyExchangeHandler{
			E: e,
		}
		handler := http.HandlerFunc(c.ServeHTTP)

		handler.ServeHTTP(rr, req)
		Expect(rr.Code).To(Equal(200))

		var body struct{
			Amount int
		}
		_ = json.Unmarshal(rr.Body.Bytes(), &body)
		Expect(body.Amount).To(Equal(300))

		close(done)
	})


})
