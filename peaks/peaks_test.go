package peaks_test

import (
	"github.com/jmnarloch/codility-go/peaks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Peaks", func() {

	var A []int

	BeforeEach(func() {
		A = []int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}
	})

	Context("Counting peaks", func() {

		It("should return peaks count", func() {
			Expect(peaks.Solution(A)).To(Equal(3))
		})
	})
})
