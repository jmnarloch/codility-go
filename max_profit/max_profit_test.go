package max_profit

import "testing"

func Test(t *testing.T) {

	// given
	A := []int{23171, 21011, 21123, 21366, 21013, 21367}

	// when
	result := Solution(A)

	// then
	if result != 356 {
		t.Error("Incorrect result", result)
	}
}
