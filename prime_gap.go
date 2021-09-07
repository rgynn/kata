package kata

import (
	"math"
)

func Gap(g, m, n int) []int {
mainloop:
	for x := m; x <= n-g; x++ {
		if IsPrime(x) {
			for y := 1; y <= g; y++ {
				if IsPrime(x + y) {
					if y == g {
						return []int{x, x + y}
					} else {
						continue mainloop
					}
				}
			}
		}
	}
	return nil
}

func IsPrime(n int) bool {
	for i := 2; i < int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
