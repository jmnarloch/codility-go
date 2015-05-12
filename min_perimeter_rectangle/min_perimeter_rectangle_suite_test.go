package min_perimeter_rectangle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMinPerimeterRectangle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MinPerimeterRectangle Suite")
}
