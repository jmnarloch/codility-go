package fish_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/fish"
)

var _ = Describe("Fish", func() {

	var A []int
	var B []int

	BeforeEach(func() {
		A = []int{4, 3, 2, 1, 5}
		B = []int{0, 1, 0, 0, 0}
	})

	Context("Counting survived fish", func() {

		It("should return live fish", func() {
			Expect(fish.Solution(A, B)).To(Equal(2))
		})
	})
})
