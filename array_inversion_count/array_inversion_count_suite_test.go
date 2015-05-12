package array_inversion_count_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestArrayInversionCount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ArrayInversionCount Suite")
}
