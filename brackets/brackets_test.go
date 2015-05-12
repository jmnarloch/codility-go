package brackets_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/brackets"
)

var _ = Describe("Brackets", func() {

	// given
	S := "{[()()]}"

	Context("Brackets nesting", func() {

		It("should valid correctness", func() {
			Expect(brackets.Solution(S)).To(Equal(1))
		})
	})
})
