package abs_distinct
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{-5, -3, -1, 0, 3, 6}

	// when
	result := Solution(A)

	// then
	if result != 5 {
		t.Error("Incorrect result", result)
	}
}