package nesting_test

import (
	"github.com/jmnarloch/codility-go/nesting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Nesting", func() {

	S := "(()(())())"

	Context("Validating paratheneses", func() {

		It("should return valid", func() {
			Expect(nesting.Solution(S)).To(Equal(1))
		})
	})
})
