package passing_cars_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPassingCars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PassingCars Suite")
}
