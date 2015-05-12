package min_perimeter_rectangle

func Solution(N int) int {
	min := (N + 1) * 2
	for div := 1; div*div <= N; div++ {
		if N%div == 0 {
			min = Min(min, (div+(N/div))*2)
		}
	}
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
