package tape_equilibrium_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/tape_equilibrium"
)

var _ = Describe("TapeEquilibrium", func() {

	var A []int

	BeforeEach(func() {
		A = []int{3, 1, 2, 4, 3}
	})

	Context("Tape", func() {

		It("should return minimal absolute diff", func() {
			Expect(tape_equilibrium.Solution(A)).To(Equal(1))
		})
	})
})
