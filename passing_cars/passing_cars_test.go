package passing_cars_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/passing_cars"
)

var _ = Describe("PassingCars", func() {

	var A []int

	BeforeEach(func() {
		A = []int{0, 1, 0, 1, 1}
	})

	Context("Passing cars", func() {

		It("should count cars", func() {
			Expect(passing_cars.Solution(A)).To(Equal(5))
		})
	})
})
