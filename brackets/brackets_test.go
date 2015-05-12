package brackets

import "testing"

func Test(t *testing.T) {

	// given
	S := "{[()()]}"

	// when
	result := Solution(S)

	// then
	if result != 1 {
		t.Error("Incorrect result", result)
	}
}
