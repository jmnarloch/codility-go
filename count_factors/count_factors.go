package count_factors

func Solution(N int) (factors int) {
	div := 1
	factors = 0

	for div*div < N {
		if N%div == 0 {
			factors += 2
		}
		div++
	}

	if div*div == N {
		factors++
	}
	return factors
}
