package max_counters

import "testing"

func Test(t *testing.T) {

	// given
	N := 5
	A := []int{3, 4, 4, 6, 1, 4, 4}

	// when
	result := Solution(N, A)

	// then
	if !Equals([]int{3, 2, 2, 4, 2}, result) {
		t.Error("Incorrect result", result)
	}
}

func Equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for ind := range a {

		if a[ind] != b[ind] {
			return false
		}
	}
	return true
}
