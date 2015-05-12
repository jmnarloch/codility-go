package tree_height_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTreeHeight(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TreeHeight Suite")
}
