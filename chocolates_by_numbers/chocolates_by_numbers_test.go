package chocolates_by_numbers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/chocolates_by_numbers"
)

var _ = Describe("ChocolatesByNumbers", func() {

	N := 10
	M := 4

	Context("Counting chocolates", func() {

		It("should return cycle size", func() {
			Expect(chocolates_by_numbers.Solution(N, M)).To(Equal(5))
		})
	})
})
