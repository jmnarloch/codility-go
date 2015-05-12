package max_slice_sum

func Solution(A []int) int {
	maxSlice := A[0]
	max := A[0]

	for ind := 1; ind < len(A); ind++ {

		maxSlice = Max(A[ind], maxSlice+A[ind])
		max = Max(max, maxSlice)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
