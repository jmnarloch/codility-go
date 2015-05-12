package fish

func Solution(A []int, B []int) int {
	N := len(A)

	if N == 0 {
		return 0
	}

	stack := make([]int, N)
	top := 0

	live := 0
	for ind := 0; ind < N; ind++ {

		if B[ind] == 0 {
			for top > 0 && stack[top-1] < A[ind] {
				top--
			}
			if top == 0 {
				live++
			}
		} else {
			stack[top] = A[ind]
			top++
		}
	}

	return live + top
}
