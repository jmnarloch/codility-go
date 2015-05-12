package dominator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDominator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dominator Suite")
}
