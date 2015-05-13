package count_distinct_slices_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCountDistinctSlices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountDistinctSlices Suite")
}
