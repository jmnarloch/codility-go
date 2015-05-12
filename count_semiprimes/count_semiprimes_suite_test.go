package count_semiprimes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCountSemiprimes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountSemiprimes Suite")
}
