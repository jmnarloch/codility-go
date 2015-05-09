package brackets

func Solution(S string) int {
	// write your code in Go 1.4

	N := len(S)
	stack := make([]rune, N)
	top := 0

	brackets := map[rune]rune {
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, c := range S {

		if _, ok := brackets[c]; ok {
			stack[top] = c
			top++
		} else if top == 0 || brackets[stack[top - 1]] != c {
			return 0
		} else {
			top--
		}
	}

	if top == 0 {
		return 1
	} else {
		return 0
	}
}