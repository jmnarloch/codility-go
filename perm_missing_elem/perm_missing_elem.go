package perm_missing_elem

func Solution(A []int) int {
	N := len(A)
	sum := 0
	var expectedSum int = (N + 1) * (N + 2) / 2
	for _, val := range A {
		sum += val
	}

	return expectedSum - sum
}
