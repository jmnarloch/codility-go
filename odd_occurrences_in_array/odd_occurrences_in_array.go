package odd_occurrences_in_array

func Solution(A []int) int {
	single := 0
	for _, val := range A {

		single ^= val
	}
	return single
}
