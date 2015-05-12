package abs_distinct_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAbsDistinct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AbsDistinct Suite")
}
