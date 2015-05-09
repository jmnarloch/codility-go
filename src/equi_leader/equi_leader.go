package equi_leader

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)

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

	count := 0
	for _, val := range A {
		if val == candidate {
			count++
		}
	}

	if count <= N / 2 {
		return 0
	}

	matching := 0
	result := 0
	for ind, val := range A {

		if val == candidate {
			matching++
		}

		if matching > (ind + 1) / 2 &&
			(count - matching > (N - ind - 1) / 2) {
			result++
		}
	}

	return result
}