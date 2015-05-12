package missing_integer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/missing_integer"
)

var _ = Describe("MissingInteger", func() {

	var A []int

	BeforeEach(func() {
		A = []int{1, 3, 6, 4, 1, 2}
	})

	Context("Founding missing possitive integer inversions", func() {

		It("should return missing integer", func() {
			Expect(missing_integer.Solution(A)).To(Equal(5))
		})
	})
})
