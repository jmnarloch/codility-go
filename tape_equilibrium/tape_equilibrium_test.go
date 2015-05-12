package tape_equilibrium

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{3, 1, 2, 4, 3}

	// when
	result := Solution(A)

	// then
	if result != 1 {
		t.Error("Incorrect result", result)
	}
}
