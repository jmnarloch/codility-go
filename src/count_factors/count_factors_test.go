package count_factors
import "testing"

func Test(t *testing.T) {

	// given
	N := 24

	// when
	result := Solution(N)

	// then
	if result != 8 {
		t.Error("Incorrect result", result)
	}
}
