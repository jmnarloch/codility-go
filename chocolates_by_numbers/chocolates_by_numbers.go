package chocolates_by_numbers

func Solution(N int, M int) int {
	return lcm(N, M) / M
}

func lcm(a int, b int) int {

	return a * b / gcd(a, b)
}

func gcd(a int, b int) int {

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
