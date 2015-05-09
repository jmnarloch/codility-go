package dominator

func Solution(A []int) int {
	// write your code in Go 1.4

	if len(A) == 0 {
		return -1
	}

	candidate := A[0]
	occur := 1
	for _, val := range A {

		if val == candidate {
			occur++
		} else {
			occur--
		}

		if occur == 0 {
			candidate = val
			occur = 1
		}
	}

	index := -1
	count := 0
	for ind, val := range A {
		if val == candidate {
			if index == -1 {
				index = ind
			}

			count++
		}
	}

	if count <= len(A) / 2 {
		return -1
	}
	return index
}