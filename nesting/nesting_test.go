package nesting_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/nesting"
)

var _ = Describe("Nesting", func() {

	S := "(()(())())"

	Context("Validating paratheneses", func() {

		It("should return valid", func() {
			Expect(nesting.Solution(S)).To(Equal(1))
		})
	})
})
