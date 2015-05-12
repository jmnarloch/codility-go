package tie_ropes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTieRopes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TieRopes Suite")
}
