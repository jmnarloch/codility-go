package max_counters_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMaxCounters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxCounters Suite")
}
