package array_inversion_count

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	aux := make([]int, N)
	return Inversions(A, aux, 0, N - 1)
}

func Inversions(A []int, aux []int, lo int, hi int) int {

	if (hi <= lo) {
		return 0
	}

	mid := lo + (hi - lo) / 2
	inv := Inversions(A, aux, lo, mid)
	inv += Inversions(A, aux, mid + 1, hi)
	inv += Merge(A, aux, lo, mid, hi)

	if inv > 1E9 {
		return -1
	}
	return inv
}

func Merge(A []int, aux []int, lo int, mid int, hi int) int {

	inv := 0
	for k := lo; k <= hi; k++ {
		aux[k] = A[k]
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {

		if i > mid {
			A[k] = aux[j]
			j++
		} else if j > hi {
			A[k] = aux[i]
			i++
		} else if less(aux[j], aux[i]) {
			inv += mid - i + 1

			A[k] = aux[j]
			j++
		} else {
			A[k] = aux[i]
			i++
		}
	}
	return inv
}

func less(v, w int) bool {
	return v < w
}