package frog_river_one
import "testing"

func Test(t *testing.T) {

	// given
	X := 5
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}

	// when
	result := Solution(X, A)

	// then
	if result != 6 {
		t.Error("Incorrect result", result)
	}
}
