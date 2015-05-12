package odd_occurrences_in_array_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOddOccurrencesInArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OddOccurrencesInArray Suite")
}
