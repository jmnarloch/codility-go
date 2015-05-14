package min_abs_sum_of_two

import (
	"sort"
)

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	aux := make([]int64, N)
	sort.Ints(A)

	lo := 0
	hi := N - 1

	for ind := N - 1; ind >= 0; ind-- {

		if abs(A[hi]) > abs(A[lo]) {
			aux[ind] = int64(A[hi])
			hi--
		} else {
			aux[ind] = int64(A[lo])
			lo++
		}
	}

	minSum := abs64(aux[0] << 1)
	for ind := 1; ind < N; ind++ {

		minSum = min64(minSum, abs64(aux[ind]<<1))
		minSum = min64(minSum, abs64(aux[ind]+aux[ind-1]))
	}

	return int(minSum)
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func abs64(n int64) int64 {
	if n < 0 {
		n = -n
	}
	return n
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
