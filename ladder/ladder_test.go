package ladder_test

import (
	"github.com/jmnarloch/codility-go/ladder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ladder", func() {

	var A []int
	var B []int

	BeforeEach(func() {
		A = []int{4, 4, 5, 5, 1}
		B = []int{3, 2, 4, 3, 1}
	})

	Context("Climbing ladder", func() {

		It("should return the number possible steps", func() {
			Expect(ladder.Solution(A, B)).To(Equal([]int{5, 1, 8, 0, 1}))
		})
	})
})
