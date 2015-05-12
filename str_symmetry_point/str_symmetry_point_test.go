package str_symmetry_point_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/codility-go/str_symmetry_point"
)

var _ = Describe("StrSymetryPoint", func() {

	S := "racecar"

	Context("Symetry point", func() {

		It("should return string symetry index", func() {
			Expect(str_symmetry_point.Solution(S)).To(Equal(3))
		})
	})
})
