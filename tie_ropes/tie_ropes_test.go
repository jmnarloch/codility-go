package tie_ropes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/tie_ropes"
)

var _ = Describe("TieRopes", func() {

	K := 4
	var A []int

	BeforeEach(func() {
		A = []int{1, 2, 3, 4, 1, 1, 3}
	})

	Context("Tie ropes", func() {

		It("should gready tie ropes", func() {
			Expect(tie_ropes.Solution(K, A)).To(Equal(3))
		})
	})
})
