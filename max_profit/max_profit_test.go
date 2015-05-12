package max_profit_test

import (
	"github.com/jmnarloch/codility-go/max_profit"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxProfit", func() {

	var A []int

	BeforeEach(func() {
		A = []int{23171, 21011, 21123, 21366, 21013, 21367}
	})

	Context("Maximum profit", func() {

		It("should return maximum profit", func() {
			Expect(max_profit.Solution(A)).To(Equal(356))
		})
	})
})
