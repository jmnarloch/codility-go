package count_triangles_test

import (
	"github.com/jmnarloch/codility-go/count_triangles"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountTriangles", func() {

	var A []int

	BeforeEach(func() {
		A = []int{10, 2, 5, 1, 8, 12}
	})

	Context("Counting triangles", func() {

		It("should return triangles count", func() {
			Expect(count_triangles.Solution(A)).To(Equal(4))
		})
	})
})
