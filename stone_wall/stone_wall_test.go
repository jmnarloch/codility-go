package stone_wall

import "testing"

func Test(t *testing.T) {

	// given
	H := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}

	// when
	result := Solution(H)

	// then
	if result != 7 {
		t.Error("Incorrect result", result)
	}
}
