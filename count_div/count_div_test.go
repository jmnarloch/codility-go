package count_div_test

import (
	"github.com/jmnarloch/codility-go/count_div"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountDiv", func() {

	A := 6
	B := 11
	K := 2

	Context("Counting divisors", func() {

		It("should return the divisors count", func() {
			Expect(count_div.Solution(A, B, K)).To(Equal(3))
		})
	})
})
