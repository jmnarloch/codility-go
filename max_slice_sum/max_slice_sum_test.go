package max_slice_sum

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{3, 2, -6, 4, 0}

	// when
	result := Solution(A)

	// then
	if result != 5 {
		t.Error("Incorrect result", result)
	}
}
