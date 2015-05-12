package perm_missing_elem_test

import (
	"github.com/jmnarloch/codility-go/perm_missing_elem"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PermMissingElem", func() {

	var A []int

	BeforeEach(func() {
		A = []int{2, 3, 1, 5}
	})

	Context("Permutation", func() {

		It("should return missing element", func() {
			Expect(perm_missing_elem.Solution(A)).To(Equal(4))
		})
	})
})
