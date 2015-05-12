package stone_wall_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStoneWall(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StoneWall Suite")
}
