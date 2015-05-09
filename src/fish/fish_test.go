package fish
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{4, 3, 2, 1, 5}
	B := []int{0, 1, 0, 0, 0}

	// when
	result := Solution(A, B)

	// then
	if result != 2 {
		t.Error("Incorrect result", result)
	}
}