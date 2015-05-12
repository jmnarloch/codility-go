package count_factors_test

import (
	"github.com/jmnarloch/codility-go/count_factors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountFactors", func() {

	N := 24

	Context("Counting factors", func() {

		It("should return factors count", func() {
			Expect(count_factors.Solution(N)).To(Equal(8))
		})
	})
})
