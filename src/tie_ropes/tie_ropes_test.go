package tie_ropes
import "testing"

func Test(t *testing.T) {

	// given
	K := 4
	A := []int{1, 2, 3, 4, 1, 1, 3}

	// when
	result := Solution(K, A)

	// then
	if result != 3 {
		t.Error("Incorrect result", result)
	}
}
