package max_nonoverlapping_segments_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMaxNonoverlappingSegments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxNonoverlappingSegments Suite")
}
