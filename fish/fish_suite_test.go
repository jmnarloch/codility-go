package fish_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFish(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fish Suite")
}
