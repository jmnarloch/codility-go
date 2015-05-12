package min_perimeter_rectangle_test

import (
	"github.com/jmnarloch/codility-go/min_perimeter_rectangle"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MinPerimeterRectangle", func() {

	N := 30

	Context("Computing perimeter", func() {

		It("should return min", func() {
			Expect(min_perimeter_rectangle.Solution(N)).To(Equal(22))
		})
	})
})
