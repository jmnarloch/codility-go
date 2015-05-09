package max_product_of_three
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{-3, 1, 2, -2, 5, 6}

	// when
	result := Solution(A)

	// then
	if result != 60 {
		t.Error("Incorrect result", result)
	}
}
