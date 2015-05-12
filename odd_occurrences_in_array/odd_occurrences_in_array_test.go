package odd_occurrences_in_array_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/odd_occurrences_in_array"
)

var _ = Describe("OddOccurencesInArray", func() {
	
	var A []int

	BeforeEach(func() {
		A = []int{9, 3, 9, 3, 9, 7, 9}
	})

	Context("Search odd occurences", func() {

		It("should find element", func() {
			Expect(odd_occurrences_in_array.Solution(A)).To(Equal(7))
		})
	})
})
