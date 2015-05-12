package max_nonoverlapping_segments

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{1, 3, 7, 9, 9}
	B := []int{5, 6, 8, 9, 10}

	// when
	result := Solution(A, B)

	// then
	if result != 3 {
		t.Error("Incorrect result", result)
	}
}
