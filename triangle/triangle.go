package triangle

import "sort"

func Solution(A []int) int {
	sort.Ints(A)

	for ind := 2; ind < len(A); ind++ {

		if A[ind-2] < A[ind-1]+A[ind] &&
			A[ind-1] < A[ind-2]+A[ind] &&
			A[ind] < A[ind-2]+A[ind-1] {

			return 1
		}
	}

	return 0
}
