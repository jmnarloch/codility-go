package peaks_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPeaks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Peaks Suite")
}
