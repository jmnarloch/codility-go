package stone_wall

func Solution(H []int) int {
	N := len(H)
	if N == 0 {
		return 0
	}

	stack := make([]int, N)
	top := 0

	blocks := 0
	for _, val := range H {

		for top > 0 && stack[top-1] > val {
			top--
		}

		if top == 0 || stack[top-1] != val {
			stack[top] = val
			top++
			blocks++
		}
	}

	return blocks
}
