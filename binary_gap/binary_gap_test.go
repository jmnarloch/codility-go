package binary_gap_test

import (
	"github.com/jmnarloch/codility-go/binary_gap"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinaryGap", func() {

	N := 1041

	Context("Counting b", func() {

		It("should return max binary gap", func() {
			Expect(binary_gap.Solution(N)).To(Equal(5))
		})
	})
})
