package perm_check_test

import (
	"github.com/jmnarloch/codility-go/perm_check"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PermCheck", func() {

	var A []int

	BeforeEach(func() {
		A = []int{4, 1, 3, 2}
	})

	Context("Permutation", func() {

		It("should return one", func() {
			Expect(perm_check.Solution(A)).To(Equal(1))
		})
	})
})
