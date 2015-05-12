package min_perimeter_rectangle

import "testing"

func Test(t *testing.T) {

	// given
	N := 30

	// when
	result := Solution(N)

	// then
	if result != 22 {
		t.Error("Incorrect result", result)
	}
}
