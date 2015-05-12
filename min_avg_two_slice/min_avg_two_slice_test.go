package min_avg_two_slice

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{4, 2, 2, 5, 1, 5, 8}

	// when
	result := Solution(A)

	// then
	if result != 1 {
		t.Error("Incorrect result", result)
	}
}
