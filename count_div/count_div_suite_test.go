package count_div_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCountDiv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountDiv Suite")
}
