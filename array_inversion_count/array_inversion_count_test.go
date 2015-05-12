package array_inversion_count_test

import (
	"github.com/jmnarloch/codility-go/array_inversion_count"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArrayInversionCount", func() {

	var A []int

	BeforeEach(func() {
		A = []int{-1, 6, 3, 4, 7, 4}
	})

	Context("Counting inversions", func() {

		It("should return inversions count", func() {
			Expect(array_inversion_count.Solution(A)).To(Equal(4))
		})
	})
})
