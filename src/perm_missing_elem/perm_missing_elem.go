package main
import "fmt"

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	sum := 0
	var expectedSum int = (N + 1) * (N + 2) / 2
	for _, val := range A {
		sum += val
	}

	return expectedSum - sum
}

func main() {

	A := []int{2, 3, 1, 5}
	fmt.Println(Solution(A))
}