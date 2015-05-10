package binary_gap
import "testing"

func Test(t *testing.T) {

	// given
	N := 1041

	// when
	result := Solution(N)

	// then
	if result != 5 {
		t.Error("Incorrect result", result)
	}
}
