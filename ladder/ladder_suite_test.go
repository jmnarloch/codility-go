package ladder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLadder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ladder Suite")
}
