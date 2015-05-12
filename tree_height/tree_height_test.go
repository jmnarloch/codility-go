package tree_height_test

import (
	. "github.com/jmnarloch/codility-go/tree_height"

	"github.com/jmnarloch/codility-go/tree_height"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TreeHeight", func() {

	var root *tree_height.Tree

	BeforeEach(func() {
		root = &Tree{X: 5, L: &Tree{X: 3, L: &Tree{X: 20}, R: &Tree{X: 21}}, R: &Tree{X: 10, L: &Tree{X: 1}}}
	})

	Context("Binary tree", func() {

		It("should return tree height", func() {
			Expect(tree_height.Solution(root)).To(Equal(2))
		})
	})
})
