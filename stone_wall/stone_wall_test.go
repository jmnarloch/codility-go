package stone_wall_test

import (
	"github.com/jmnarloch/codility-go/stone_wall"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StoneWall", func() {

	var A []int

	BeforeEach(func() {
		A = []int{8, 8, 5, 7, 9, 8, 7, 4, 8}
	})

	Context("Stone wall", func() {

		It("should return minimum number of blocks", func() {
			Expect(stone_wall.Solution(A)).To(Equal(7))
		})
	})
})
