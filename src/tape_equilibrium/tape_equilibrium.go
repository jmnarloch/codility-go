package main
import (
	"fmt"
)

func Solution(A []int) (min int) {
	// write your code in Go 1.4
	total := 0

	for _, val := range A {
		total += val
	}

	left := total - A[len(A) - 1]
	right := A[len(A) - 1]
	min = Abs(left - right)
	for ind := len(A) - 2; ind >= 1; ind-- {
		left -= A[ind]
		right += A[ind]

		min = Min(min, Abs(left - right))
	}
	return
}

func Abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	A := []int{3, 1, 2, 4, 3}
	fmt.Println(Solution(A))
}
