package odd_occurrences_in_array

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{9, 3, 9, 3, 9, 7, 9}

	// when
	result := Solution(A)

	// then
	if result != 7 {
		t.Error("Incorrect result", result)
	}
}
