package min_abs_sum_of_two_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMinAbsSumOfTwo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MinAbsSumOfTwo Suite")
}
