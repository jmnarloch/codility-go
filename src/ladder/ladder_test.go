package ladder
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{4, 4, 5, 5, 1}
	B := []int{3, 2, 4, 3, 1}

	// when
	result := Solution(A, B)

	// then
	if !Equals([]int{5, 1, 8, 0, 1}, result) {
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