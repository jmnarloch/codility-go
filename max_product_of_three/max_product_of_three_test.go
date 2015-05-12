package max_product_of_three_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/max_product_of_three"
)

var _ = Describe("MaxProductOfThree", func() {

	var A []int

	BeforeEach(func() {
		A = []int{-3, 1, 2, -2, 5, 6}
	})

	Context("Max product", func() {

		It("should return max product of three elements", func() {
			Expect(max_product_of_three.Solution(A)).To(Equal(60))
		})
	})
})
