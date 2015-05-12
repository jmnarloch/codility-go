package perm_check

func Solution(A []int) int {
	N := len(A)

	counts := make([]int, N)
	for _, val := range A {
		if val < 1 || val > N {
			return 0
		}
		counts[val-1]++
	}

	for _, occ := range counts {
		if occ != 1 {
			return 0
		}
	}

	return 1
}
