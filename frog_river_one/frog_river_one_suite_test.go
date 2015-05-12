package frog_river_one_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFrogRiverOne(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FrogRiverOne Suite")
}
