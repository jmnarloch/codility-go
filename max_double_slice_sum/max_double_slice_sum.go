package max_double_slice_sum

func Solution(A []int) int {
	N := len(A)
	left := make([]int, N)
	right := make([]int, N)

	leftMaxSlice := 0
	rightMaxSlice := 0
	for ind := 1; ind < N-1; ind++ {
		leftMaxSlice = max(0, A[ind]+leftMaxSlice)
		left[ind] = leftMaxSlice

		rightMaxSlice = max(0, A[N-ind-1]+rightMaxSlice)
		right[N-ind-1] = rightMaxSlice
	}

	maxResult := 0
	for ind := 0; ind < N-2; ind++ {
		maxResult = max(maxResult, left[ind]+right[ind+2])
	}
	return maxResult
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
