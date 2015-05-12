package tree_height

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func Solution(T *Tree) int {
	if T == nil {
		return -1
	}

	return Max(Solution(T.L), Solution(T.R)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
