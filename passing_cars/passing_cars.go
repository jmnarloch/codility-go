package passing_cars

func Solution(A []int) (cars int) {
	east := 0
	cars = 0

	for _, val := range A {

		if val == 0 {
			east++
		} else {
			cars += east
		}

		if cars > 1000000000 {
			return -1
		}
	}
	return
}
