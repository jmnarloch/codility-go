package max_profit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMaxProfit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxProfit Suite")
}
