package main

import "fmt"

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)

	counts := make([]int, N)
	for _, val := range(A) {
		if val < 1 || val > N {
			return 0
		}
		counts[val - 1]++
	}

	for _, occ := range(counts) {
		if occ != 1 {
			return 0
		}
	}

	return 1
}

func main() {
	A := []int{4, 1, 3, 2}
	fmt.Println(Solution(A))
}
