package main
import "sort"

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)

	sort.Ints(A)

	return Max(A[0] * A[1] * A[N - 1], A[N - 3] * A[N - 2] * A[N - 1])
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
