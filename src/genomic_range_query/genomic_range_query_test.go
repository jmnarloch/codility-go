package genomic_range_query
import (
	"testing"
)

func Test(t *testing.T) {

	// given
	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}

	// when
	result := Solution(S, P, Q)

	// then
	if !Equals([]int{2, 4, 1}, result)  {
		t.Error("Incorrect result", result)
	}
}

func Equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for ind := range a {

		if a[ind] != b[ind] {
			return false
		}
	}
	return true
}
