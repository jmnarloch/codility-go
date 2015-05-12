package genomic_range_query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGenomicRangeQuery(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GenomicRangeQuery Suite")
}
