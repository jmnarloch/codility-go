package frog_river_one

func Solution(X int, A []int) (min int) {
	leaves := make([]int, X)
	for i := range leaves {
		leaves[i] = -1
	}

	for ind, val := range A {
		if val <= X && leaves[val-1] == -1 {
			leaves[val-1] = ind
		}
	}

	min = leaves[0]
	for _, val := range leaves {
		if val == -1 {
			return -1
		}
		min = Max(min, val)
	}

	return
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
