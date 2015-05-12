package chocolates_by_numbers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChocolatesByNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ChocolatesByNumbers Suite")
}
