package perm_missing_elem_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPermMissingElem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PermMissingElem Suite")
}
