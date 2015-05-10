package count_factors

func Solution(N int) (factors int) {
	// write your code in Go 1.4

	div := 1
	factors = 0

	for div * div < N {
		if N % div == 0 {
			factors += 2
		}
		div++
	}

	if div * div == N {
		factors++
	}
	return factors
}
