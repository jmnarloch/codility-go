package triangle

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{10, 2, 5, 1, 8, 20}

	// when
	result := Solution(A)

	// then
	if result != 1 {
		t.Error("Incorrect result", result)
	}
}
