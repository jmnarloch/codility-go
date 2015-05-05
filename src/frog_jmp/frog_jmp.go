package main

import (
	"fmt"
	"math"
)

func Solution(X int, Y int, D int) int {
	return int(math.Ceil(float64(Y - X) / float64(D)))
}

func main() {
	X := 10
	Y := 85
	D := 30

	fmt.Println(Solution(X, Y, D))
}