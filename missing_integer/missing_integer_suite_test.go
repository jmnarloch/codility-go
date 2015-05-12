package missing_integer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMissingInteger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MissingInteger Suite")
}
