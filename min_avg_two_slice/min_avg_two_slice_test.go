package min_avg_two_slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/min_avg_two_slice"
)

var _ = Describe("MinAvgTwoSlice", func() {

	var A []int

	BeforeEach(func() {
		A = []int{4, 2, 2, 5, 1, 5, 8}
	})

	Context("Minimum average of slice", func() {

		It("should return min avg slice", func() {
			Expect(min_avg_two_slice.Solution(A)).To(Equal(1))
		})
	})
})
