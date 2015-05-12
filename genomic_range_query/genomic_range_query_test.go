package genomic_range_query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/genomic_range_query"
)

var _ = Describe("GenomicRangeQuery", func() {

	S := "CAGCCTA"
	var P []int
	var Q []int

	BeforeEach(func() {
		P = []int{2, 5, 0}
		Q = []int{4, 5, 6}
	})

	Context("Genoms", func() {

		It("should return range query result", func() {
			Expect(genomic_range_query.Solution(S, P, Q)).To(Equal([]int{2, 4, 1}))
		})
	})
})
