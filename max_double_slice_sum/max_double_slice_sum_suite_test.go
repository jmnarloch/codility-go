package max_double_slice_sum_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMaxDoubleSliceSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxDoubleSliceSum Suite")
}
