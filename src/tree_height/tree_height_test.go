package tree_height
import "testing"

func Test(t *testing.T) {

	// given
	root := &Tree{X:5, L:&Tree{X:3, L:&Tree{X: 20}, R:&Tree{X: 21}}, R: &Tree{X: 10, L: &Tree{X: 1}}}

	// when
	result := Solution(root)

	// then
	if result != 2 {
		t.Error("Incorrect result", result)
	}
}
