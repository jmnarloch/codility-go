package distinct

import "sort"

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	distinct := 1
	for ind := 1; ind < len(A); ind++ {
		if A[ind] == A[ind-1] {
			continue
		}
		distinct++
	}
	return distinct
}
