package count_div
import "fmt"

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4

	div := (B / K - A / K);
	if A % K == 0 {
		div++
	}
	return div
}

func main() {

	A := 6
	B := 11
	K := 2
	fmt.Println(Solution(A, B, K))
}