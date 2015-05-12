package max_product_of_three

import "sort"

func Solution(A []int) int {
	N := len(A)
	sort.Ints(A)

	return max(A[0]*A[1]*A[N-1], A[N-3]*A[N-2]*A[N-1])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
