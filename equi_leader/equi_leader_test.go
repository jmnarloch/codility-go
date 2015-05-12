package equi_leader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/equi_leader"
)

var _ = Describe("EquiLeader", func() {

	var A []int

	BeforeEach(func() {
		A = []int{4, 3, 4, 4, 4, 2}
	})

	Context("Searching leader", func() {

		It("should return equi leader counts", func() {
			Expect(equi_leader.Solution(A)).To(Equal(2))
		})
	})
})
