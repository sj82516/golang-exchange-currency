package src

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCurrencyExchange(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Currency Exchange")
}
