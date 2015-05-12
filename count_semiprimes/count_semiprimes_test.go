package count_semiprimes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/count_semiprimes"
)

var _ = Describe("CountSemiprimes", func() {

	N := 26
	var P []int
	var Q []int

	BeforeEach(func() {
		P = []int{1, 4, 16}
		Q = []int{26, 10, 20}
	})

	Context("Counting semi primes", func() {

		It("should return count of semi primes", func() {
			Expect(count_semiprimes.Solution(N, P, Q)).To(Equal([]int{10, 4, 0}))
		})
	})
})
