package peaks

func Solution(A []int) int {

	N := len(A)
	peaks := make([]int, N)

	for ind := range A {
		if ind > 0 {
			peaks[ind] = peaks[ind-1]
		}

		if ind > 0 && ind < N-1 && A[ind] > A[ind-1] && A[ind] > A[ind+1] {
			peaks[ind]++
		}
	}

	if peaks[N-1] == 0 {
		return 0
	}

	div := 1
	blocks := 0

	for div*div <= N {

		if N%div == 0 {
			blockSize := N / div
			hasPeak := true

			for ind := N - 1; ind > blockSize-1; ind -= blockSize {
				if peaks[ind] == peaks[ind-blockSize] {
					hasPeak = false
					break
				}
			}

			if hasPeak {
				blocks = max(blocks, N/blockSize)
			}

			blockSize = div
			hasPeak = true
			for ind := N - 1; ind > blockSize-1; ind -= blockSize {
				if peaks[ind] == peaks[ind-blockSize] {
					hasPeak = false
					break
				}
			}

			if hasPeak {
				blocks = max(blocks, N/div)
			}
		}
		div++
	}

	return blocks
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
