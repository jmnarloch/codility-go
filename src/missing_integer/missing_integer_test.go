package missing_integer
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{1, 3, 6, 4, 1, 2}

	// when
	result := Solution(A)

	// then
	if result != 5 {
		t.Error("Incorrect result", result);
	}
}