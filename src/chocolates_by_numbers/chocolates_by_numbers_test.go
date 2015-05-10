package chocolates_by_numbers
import "testing"

func Test(t *testing.T) {

	// given
	N := 10
	M := 4

	// when
	result := Solution(N, M)

	// then
	if result != 5 {
		t.Error("Incorrect result", result)
	}
}
