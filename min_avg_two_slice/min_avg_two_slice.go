package min_avg_two_slice

func Solution(A []int) int {
	minAvg := float64(A[0]+A[1]) / 2.0
	var index int = 0

	for ind := 1; ind < len(A); ind++ {

		avg := float64(A[ind]+A[ind-1]) / 2.0
		if avg < minAvg {
			index = ind - 1
			minAvg = avg
		}

		if ind >= 2 {
			avg = float64(A[ind]+A[ind-1]+A[ind-2]) / 3.0
			if avg < minAvg {
				index = ind - 2
				minAvg = avg
			}
		}
	}

	return index
}
