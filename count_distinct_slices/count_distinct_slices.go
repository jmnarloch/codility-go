package count_distinct_slices

func Solution(M int, A []int) int {

	N := len(A)
	occur := make([]int, M+1)
	fill(occur, -1)

	lo := -1
	slices := 0

	for hi := 0; hi < N; hi++ {
		if occur[A[hi]] > lo {
			lo = occur[A[hi]]
		}

		slices += hi - lo
		if slices > 1E9 {
			return 1E9
		}

		occur[A[hi]] = hi
	}

	return slices
}

func fill(a []int, val int) {
	for ind := range a {
		a[ind] = val
	}
}
