package count_factors_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCountFactors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountFactors Suite")
}
