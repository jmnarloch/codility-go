package max_nonoverlapping_segments_test

import (
	"github.com/jmnarloch/codility-go/max_nonoverlapping_segments"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxNonOverlappingSeqments", func() {

	var A []int
	var B []int

	BeforeEach(func() {
		A = []int{1, 3, 7, 9, 9}
		B = []int{5, 6, 8, 9, 10}
	})

	Context("Non overlapping seqments", func() {

		It("should return non overlapping count", func() {
			Expect(max_nonoverlapping_segments.Solution(A, B)).To(Equal(3))
		})
	})
})
