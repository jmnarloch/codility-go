package count_triangles_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCountTriangles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountTriangles Suite")
}
