package odd_occurrences_in_array

func Solution(A []int) int {
	// write your code in Go 1.4

	single := 0
	for _, val := range A {

		single ^= val
	}
	return single
}
