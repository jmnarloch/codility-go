package distinct_test

import (
	"github.com/jmnarloch/codility-go/distinct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Distinct", func() {

	var A []int

	BeforeEach(func() {
		A = []int{2, 1, 1, 2, 3, 1}
	})

	Context("Counting distinct", func() {

		It("should return distinct element count", func() {
			Expect(distinct.Solution(A)).To(Equal(3))
		})
	})
})
