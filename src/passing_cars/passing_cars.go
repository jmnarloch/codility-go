package main
import "fmt"

func Solution(A []int) (cars int) {
	// write your code in Go 1.4

	east := 0
	cars = 0

	for _, val := range A {

		if val == 0 {
			east++
		} else {
			cars += east
		}

		if cars >  1000000000 {
			return -1
		}
	}
	return
}

func main() {
	A := []int{0, 1, 0, 1, 1}
	fmt.Println(Solution(A))
}
