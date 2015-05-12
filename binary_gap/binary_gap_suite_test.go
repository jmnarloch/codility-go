package binary_gap_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBinaryGap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinaryGap Suite")
}
