package frog_river_one_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/frog_river_one"
)

var _ = Describe("FrogRiverOne", func() {

	X := 5
	var A []int

	BeforeEach(func() {
		A = []int{1, 3, 1, 4, 2, 3, 5, 4}
	})

	Context("Counting minimal time to cross river", func() {

		It("should return minmal time", func() {
			Expect(frog_river_one.Solution(X, A)).To(Equal(6))
		})
	})
})
