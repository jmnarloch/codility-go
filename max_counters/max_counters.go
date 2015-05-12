package max_counters

func Solution(N int, A []int) []int {
	counters := make([]int, N)

	maxCounter := 0
	maxResult := 0

	for _, val := range A {
		if val == N+1 {
			maxResult = maxCounter
		} else {
			if counters[val-1] < maxResult {
				counters[val-1] = maxResult
			}
			counters[val-1]++
			maxCounter = max(maxCounter, counters[val-1])
		}
	}

	for ind := range counters {
		if counters[ind] < maxResult {
			counters[ind] = maxResult
		}
	}

	return counters
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
