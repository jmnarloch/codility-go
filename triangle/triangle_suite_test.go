package triangle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTriangle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Triangle Suite")
}
