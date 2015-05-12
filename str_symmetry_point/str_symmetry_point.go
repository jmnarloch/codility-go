package str_symmetry_point

func Solution(S string) int {
	N := len(S)
	if N <= 1 {
		return N - 1
	} else if N%2 == 0 {
		return -1
	}

	lo := 0
	hi := N - 1

	for lo < hi {
		if S[lo] != S[hi] {
			return -1
		}
		lo++
		hi--
	}
	return lo
}
