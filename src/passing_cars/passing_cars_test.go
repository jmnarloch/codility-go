package passing_cars
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{0, 1, 0, 1, 1}

	// when
	result := Solution(A)

	// then
	if result != 5 {
		t.Error("Incorrect result", result)
	}
}
