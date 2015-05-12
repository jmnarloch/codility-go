package nesting

func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}

	valid := 0
	for _, c := range S {

		if c == '(' {
			valid++
		} else {
			valid--
		}

		if valid < 0 {
			return 0
		}
	}

	if valid == 0 {
		return 1
	} else {
		return 0
	}
}
