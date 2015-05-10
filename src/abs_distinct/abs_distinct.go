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

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}