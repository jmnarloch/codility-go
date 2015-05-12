package triangle_test

import (
	"github.com/jmnarloch/codility-go/triangle"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Triangle", func() {

	var A []int

	BeforeEach(func() {
		A = []int{10, 2, 5, 1, 8, 20}
	})

	Context("Triangle", func() {

		It("should return one", func() {
			Expect(triangle.Solution(A)).To(Equal(1))
		})
	})
})
