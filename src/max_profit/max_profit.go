package max_profit

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	if N == 0 {
		return 0
	}

	prev := A[0]
	for ind := range A {

		tmp := A[ind]
		A[ind] -= prev
		prev = tmp
	}

	maxSlice := 0
	max := 0
	for _, val := range A {

		maxSlice = Max(0, maxSlice + val)
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