package count_distinct_slices_test

import (
	"github.com/jmnarloch/codility-go/count_distinct_slices"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountDistinctSlices", func() {

	Context("Counting unique slices", func() {

		It("should return unique slices", func() {
			M := 6
			A := []int{3, 4, 5, 5, 2}

			Expect(count_distinct_slices.Solution(M, A)).To(Equal(9))
		})

		It("should handle single element", func() {
			M := 1
			A := []int{1}

			Expect(count_distinct_slices.Solution(M, A)).To(Equal(1))
		})

		It("should handle duplicated element", func() {
			M := 1
			A := []int{1, 1}

			Expect(count_distinct_slices.Solution(M, A)).To(Equal(2))
		})

		It("should handle duplicated element", func() {

			M := 2
			A := []int{1, 2, 1}

			Expect(count_distinct_slices.Solution(M, A)).To(Equal(5))
		})
	})
})
