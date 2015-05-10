package ladder
import (
	"errors"
)

func Solution(A []int, B []int) []int {
	// write your code in Go 1.4

	N := len(A)
	fib := NewFib(N)

	result := make([]int, N)
	for ind := 0; ind < N; ind++ {

		val, _ := fib.Get(A[ind] - 1)
		result[ind] =  int(val % (uint64(1) << uint64(B[ind])))
	}

	return result
}

type Fib interface {
	Get(ind int) (uint64, error)
}

func NewFib(N int) Fib {

	seq := new(FibSeq)
	seq.init(N)
	return seq
}

type FibSeq struct {
	seq []uint64;
	size int;
}

func (fib *FibSeq) init(N int) {

	fib.size = N
	fib.seq = make([]uint64, N)

	fib.seq[0] = 1

	if N > 1 {
		fib.seq[1] = 2

		for ind := 2; ind < N; ind++ {
			fib.seq[ind] = fib.seq[ind - 1] + fib.seq[ind - 2]
		}
	}
}

func (fib *FibSeq) Get(ind int) (uint64, error) {

	if ind < 0 || ind >= fib.size {
		return 0, errors.New("Index out of bounds")
	}

	return fib.seq[ind], nil
}