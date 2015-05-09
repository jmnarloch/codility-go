package frog_jmp
import "testing"

func Test(t* testing.T) {

	// given
	X := 10
	Y := 85
	D := 30

	// when
	result := Solution(X, Y, D)

	// then
	if result != 3 {
		t.Error("Incorrect result", result)
	}
}