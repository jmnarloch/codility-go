package count_div

import "testing"

func Test(t *testing.T) {

	// given
	A := 6
	B := 11
	K := 2

	// when
	result := Solution(A, B, K)

	if result != 3 {
		t.Error("Incorrect result", result)
	}
}
