package equi_leader
import "testing"

func Test(t *testing.T) {

	// given
	A := []int{4, 3, 4, 4, 4, 2}

	// when
	result := Solution(A)

	// then
	if result != 2 {
		t.Error("Incorrect result", result)
	}
}