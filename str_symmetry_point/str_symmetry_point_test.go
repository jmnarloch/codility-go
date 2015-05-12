package str_symmetry_point

import "testing"

func Test(t *testing.T) {

	// given
	S := "racecar"

	// when
	result := Solution(S)

	// then
	if result != 3 {
		t.Error("Incorrect result", result)
	}
}
