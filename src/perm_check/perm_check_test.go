package perm_check
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{4, 1, 3, 2}

	// when
	result := Solution(A)

	// then
	if result != 1 {
		t.Error("Incorrect result", result)
	}
}