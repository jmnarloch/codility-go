package max_double_slice_sum_test

import (
	"github.com/jmnarloch/codility-go/max_double_slice_sum"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxDoubleSliceSum", func() {

	var A []int

	BeforeEach(func() {
		A = []int{3, 2, 6, -1, 4, 5, -1, 2}
	})

	Context("Counting inversions", func() {

		It("should return inversions count", func() {
			Expect(max_double_slice_sum.Solution(A)).To(Equal(17))
		})
	})
})
