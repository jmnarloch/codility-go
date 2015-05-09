package distinct
import "sort"

func Solution(A []int) int {
	// write your code in Go 1.4

	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	distinct := 1
	for ind := 1; ind < len(A); ind++ {
		if A[ind] == A[ind - 1] {
			continue
		}
		distinct++
	}
	return distinct
}
