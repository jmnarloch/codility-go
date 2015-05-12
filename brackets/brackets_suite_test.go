package brackets_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBrackets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Brackets Suite")
}
