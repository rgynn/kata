package kata

import (
	"fmt"
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	var sum int
	for i, c := range []rune(fmt.Sprintf("%d", n)) {
		d, _ := strconv.Atoi(string(c))
		sum += int(math.Pow(float64(d), float64(p+i)))
	}
	if sum%n == 0 {
		return sum / n
	}
	return -1
}
