package max_product_of_three_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMaxProductOfThree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxProductOfThree Suite")
}
