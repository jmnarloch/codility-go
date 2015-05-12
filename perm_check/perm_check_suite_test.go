package perm_check_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPermCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PermCheck Suite")
}
