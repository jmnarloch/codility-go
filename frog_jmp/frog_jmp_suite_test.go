package frog_jmp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFrogJmp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FrogJmp Suite")
}
