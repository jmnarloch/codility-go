package abs_distinct

func Solution(A []int) int {
	N := len(A)
	aux := make([]int, N)
	lo := 0
	hi := N - 1

	for k := N - 1; k >= 0; k-- {

		if abs(A[hi]) > abs(A[lo]) {
			aux[k] = abs(A[hi])
			hi--
		} else {
			aux[k] = abs(A[lo])
			lo++
		}
	}

	distinct := 1
	for ind := 1; ind < N; ind++ {
		if aux[ind] != aux[ind-1] {
			distinct++
		}
	}
	return distinct
}

func Solution2(A []int) int {
	N := len(A)
	lo := 0
	hi := N - 1

	distinct := 1
	prev := max(abs(A[lo]), abs(A[hi]))
	for lo <= hi {

		head := abs(A[lo])
		if prev == head {
			lo++
			continue
		}

		tail := abs(A[hi])
		if prev == tail {
			hi--
			continue
		}

		if head >= tail {
			prev = head
			lo++
		} else {
			prev = tail
			hi--
		}
		distinct++
	}
	return distinct
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
