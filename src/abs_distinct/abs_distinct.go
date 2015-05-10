package abs_distinct

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	aux := make([]int, N)
	lo := 0
	hi := N - 1

	for k := N - 1; k >= 0; k-- {

		if Abs(A[hi]) > Abs(A[lo]) {
			aux[k] = Abs(A[hi])
			hi--
		} else {
			aux[k] = Abs(A[lo])
			lo++
		}
	}

	distinct := 1
	for ind := 1; ind < N; ind++ {
		if aux[ind] != aux[ind - 1] {
			distinct++
		}
	}
	return distinct
}

func Solution2(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	lo := 0
	hi := N - 1

	distinct := 1
	prev := Max(Abs(A[lo]), Abs(A[hi]))
	for lo <= hi {

		head := Abs(A[lo])
		if prev == head {
			lo++
			continue
		}

		tail := Abs(A[hi])
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

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}