package binary_gap

func Solution(N int) int {
	val := uint(N)
	for val > 0 && (val&1) == 0 {
		val >>= 1
	}

	gap := 0
	maxGap := 0
	for val > 0 {
		if (val & 1) == 1 {
			maxGap = max(maxGap, gap)
			gap = 0
		} else {
			gap++
		}
		val >>= 1
	}

	return max(maxGap, gap)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
