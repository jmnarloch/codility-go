package tie_ropes

func Solution(K int, A []int) int {
	len := 0
	ropes := 0
	for _, val := range A {

		len += val

		if len >= K {
			len = 0
			ropes++
		}
	}

	return ropes
}
