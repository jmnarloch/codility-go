package max_counters

func Solution(N int, A []int) []int {
	// write your code in Go 1.4

	counters := make([]int, N)

	maxCounter := 0
	max := 0

	for _, val := range A {
		if val == N + 1 {
			max = maxCounter
		} else {
			if counters[val - 1] < max {
				counters[val - 1] = max
			}
			counters[val - 1]++
			maxCounter = Max(maxCounter, counters[val - 1])
		}
	}

	for ind := range counters {
		if counters[ind] < max {
			counters[ind] = max
		}
	}

	return counters
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
