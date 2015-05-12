package distinct

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{2, 1, 1, 2, 3, 1}

	// when
	result := Solution(A)

	// then
	if result != 3 {
		t.Error("Incorrect result", result)
	}
}
