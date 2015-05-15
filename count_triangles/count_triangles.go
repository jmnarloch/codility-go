package count_triangles

import (
	"sort"
)

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	// sort
	sort.Ints(A)

	triangles := 0
	for i := 0; i < N; i++ {
		k := i + 1
		for j := i + 1; j < N; j++ {
			for k < N && A[i]+A[j] > A[k] {
				k++
			}
			triangles += k - j - 1
		}
	}
	return triangles
}
