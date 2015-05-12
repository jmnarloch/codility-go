package count_div

func Solution(A int, B int, K int) int {
	div := (B/K - A/K)
	if A%K == 0 {
		div++
	}
	return div
}
