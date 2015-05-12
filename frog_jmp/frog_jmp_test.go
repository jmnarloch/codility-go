package frog_jmp_test

import (
	"github.com/jmnarloch/codility-go/frog_jmp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FrogJmp", func() {

	X := 10
	Y := 85
	D := 30

	Context("Counting jump", func() {

		It("should return minimum number of jump", func() {
			Expect(frog_jmp.Solution(X, Y, D)).To(Equal(3))
		})
	})
})
