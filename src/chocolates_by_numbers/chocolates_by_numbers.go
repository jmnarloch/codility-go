package chocolates_by_numbers

func Solution(N int, M int) int {
	// write your code in Go 1.4

	return Lcm(N, M) / M
}

func Lcm(a int, b int) int {

	return a * b / Gcd(a, b)
}

func Gcd(a int, b int) int {

	if b == 0 {
		return a
	}

	return Gcd(b, a % b)
}
