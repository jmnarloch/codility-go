package perm_missing_elem

import (
	"testing"
)

func Test(t *testing.T) {

	// given
	A := []int{2, 3, 1, 5}

	// when
	result := Solution(A)

	// then
	if result != 4 {
		t.Error("Incorrect reusult", result)
	}
}
