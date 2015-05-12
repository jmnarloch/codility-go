package nesting_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nesting Suite")
}
