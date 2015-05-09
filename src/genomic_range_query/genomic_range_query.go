package genomic_range_query

func Solution(S string, P []int, Q []int) []int {
	// write your code in Go 1.4

	N := len(S)
	M := len(P)
	G := 4

	genoms := Genoms()
	nucleons := make([][]int, N)

	for ind, gen := range S {
		nucleons[ind] = make([]int, G)
		for i := 0; i < G; i++ {
			nucleons[ind][i] = -1
		}

		if ind != 0 {
			for i := 0; i < G; i++ {
				nucleons[ind][i] = nucleons[ind - 1][i]
			}
		}
		nucleons[ind][genoms[gen]] = ind
	}

	result := make([]int, M)
	for ind := 0; ind < M; ind++ {
		for i := 0; i < G; i++ {
			if nucleons[Q[ind]][i] >= P[ind] {
				result[ind] = i + 1
				break
			}
		}
	}

	return result
}

func Genoms() map[rune]int {

	return map[rune]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}
}
