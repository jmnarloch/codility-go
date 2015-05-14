package min_abs_sum_of_two_test

import (
	"github.com/jmnarloch/codility-go/min_abs_sum_of_two"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MinAbsSumOfTwo", func() {

	Context("Counting minimum absolute sum", func() {

		A := []int{1, 4, -3}

		It("should return min", func() {
			Expect(min_abs_sum_of_two.Solution(A)).To(Equal(1))
		})
	})

	Context("Counting minimum absolute sum", func() {

		A := []int{-8, 4, 5, -10, 3}

		It("should return min", func() {
			Expect(min_abs_sum_of_two.Solution(A)).To(Equal(3))
		})
	})
})
