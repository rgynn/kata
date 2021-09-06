package kata

import (
	"fmt"
	"strings"
)

func NbDig(n int, d int) int {
	var result int
	for x := 0; x <= n; x++ {
		result += strings.Count(fmt.Sprintf("%d", x*x), fmt.Sprintf("%d", d))
	}
	return result
}
