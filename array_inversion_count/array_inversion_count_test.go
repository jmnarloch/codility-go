package array_inversion_count

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{-1, 6, 3, 4, 7, 4}

	// when
	result := Solution(A)

	// then
	if result != 4 {
		t.Error("Incorrect result", result)
	}
}
