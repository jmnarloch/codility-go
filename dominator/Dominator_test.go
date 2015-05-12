package dominator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/dominator"
)

var _ = Describe("Dominator", func() {

	var A []int

	BeforeEach(func() {
		A = []int{3, 4, 3, 2, 3, -1, 3, 3}
	})

	Context("Searching dominator", func() {

		It("should return any dominator index", func() {
			Expect(dominator.Solution(A)).To(Equal(0))
		})
	})
})
