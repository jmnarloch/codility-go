package distinct_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDistinct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Distinct Suite")
}
