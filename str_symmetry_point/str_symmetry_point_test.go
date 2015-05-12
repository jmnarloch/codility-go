package str_symmetry_point_test

import (
	"github.com/jmnarloch/codility-go/str_symmetry_point"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StrSymetryPoint", func() {

	S := "racecar"

	Context("Symetry point", func() {

		It("should return string symetry index", func() {
			Expect(str_symmetry_point.Solution(S)).To(Equal(3))
		})
	})
})
