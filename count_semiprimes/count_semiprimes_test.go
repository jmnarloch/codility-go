package count_semiprimes

import "testing"

func Test(t *testing.T) {

	// given
	N := 26
	P := []int{1, 4, 16}
	Q := []int{26, 10, 20}

	// when
	result := Solution(N, P, Q)

	// then
	if !Equals([]int{10, 4, 0}, result) {
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
