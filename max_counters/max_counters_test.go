package max_counters_test

import (
	"github.com/jmnarloch/codility-go/max_counters"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxCounters", func() {

	N := 5
	var A []int

	BeforeEach(func() {
		A = []int{3, 4, 4, 6, 1, 4, 4}
	})

	Context("Maximum counters", func() {

		It("should return maximum counters", func() {
			Expect(max_counters.Solution(N, A)).To(Equal([]int{3, 2, 2, 4, 2}))
		})
	})
})
