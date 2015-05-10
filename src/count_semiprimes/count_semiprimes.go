package count_semiprimes

func Solution(N int, P []int, Q []int) []int {
	// write your code in Go 1.4

	semiPrimes := CountSemiPrimes(N)

	result := make([]int, len(P))
	for ind := 0; ind < len(P); ind++ {
		result[ind] = semiPrimes[Q[ind]] - semiPrimes[P[ind] - 1]
	}

	return result
}

func EratosthenesSieve(N int) []int {

	sieve := make([]int, N + 1)
	sieve[0] = 0
	sieve[1] = 1

	for num := 2; num <= N; num++ {
		if sieve[num] == 0 {
			sieve[num] = num
			for mul := num * num; mul <= N; mul += num {
				if sieve[mul] == 0 {
					sieve[mul] = num
				}
			}
		}
	}
	return sieve
}

func CountSemiPrimes(N int) []int {

	sieve := EratosthenesSieve(N)
	semiPrimes := make([]int, N + 1)

	for ind, prim := range sieve {
		if ind > 0 {
			semiPrimes[ind] = semiPrimes[ind - 1]
		}
		if ind == prim {
			continue
		}
		if div := ind / sieve[ind]; sieve[div] == div {
			semiPrimes[ind]++
		}
	}
	return semiPrimes
}