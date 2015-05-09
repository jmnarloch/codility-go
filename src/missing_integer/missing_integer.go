package missing_integer

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)

	ind := 0
	for ind < N {
		if A[ind] != ind - 1 && A[ind] >= 1 && A[ind] <= N && A[ind] != A[A[ind] - 1] {
			Exch(A, ind, A[ind] - 1)
		} else {
			ind++
		}
	}

	for ind, val := range A {
		if ind + 1 != val {
			return ind + 1
		}
	}

	return N + 1
}

func Exch(A []int, i int, j int) {

	tmp := A[i]
	A[i] = A[j]
	A[j] = tmp
}

func Solution2(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	ints := make([]int, N)

	for _, val := range A {
		if val >= 1 && val <= N {
			ints[val - 1]++
		}
	}

	for ind, occ := range ints {
		if occ == 0 {
			return ind + 1
		}
	}

	return N + 1
}
