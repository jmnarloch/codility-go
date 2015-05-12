package max_slice_sum_test

import (
	"github.com/jmnarloch/codility-go/max_slice_sum"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxSliceSum", func() {

	var A []int

	BeforeEach(func() {
		A = []int{3, 2, -6, 4, 0}
	})

	Context("Max slice sum", func() {

		It("should return max sum", func() {
			Expect(max_slice_sum.Solution(A)).To(Equal(5))
		})
	})
})
