package equi_leader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEquiLeader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EquiLeader Suite")
}
