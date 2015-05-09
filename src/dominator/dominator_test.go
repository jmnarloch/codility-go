package dominator
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{3, 4, 3, 2, 3, -1, 3, 3}

	// when
	result := Solution(A)

	// then
	if result != 0 {
		t.Error("Incorrect result", result)
	}
}