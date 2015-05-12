package max_slice_sum_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMaxSliceSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxSliceSum Suite")
}
