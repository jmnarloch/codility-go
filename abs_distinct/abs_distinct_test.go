package abs_distinct_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/abs_distinct"
)

var _ = Describe("AbsDistinct", func() {

	var A []int

	BeforeEach(func() {
		A = []int{-5, -3, -1, 0, 3, 6}
	})

	Context("Unique elements", func() {

		It("should return distinct count", func() {
			Expect(abs_distinct.Solution2(A)).To(Equal(5))
		})
	})
})
