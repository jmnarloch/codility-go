package max_nonoverlapping_segments

func Solution(A []int, B []int) int {
	// write your code in Go 1.4

	N := len(A)
	if N <= 1 {
		return N
	}

	result := 1
	prev := B[0]
	for ind := 1; ind < N; ind++{
		if A[ind] > B[ind - 1] || A[ind] > prev {
			prev = B[ind]
			result++
		}
	}
	return result
}