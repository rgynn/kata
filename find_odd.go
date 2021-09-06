package kata

func FindOdd(seq []int) int {
	appearences := map[int]int{}
	for _, n := range seq {
		_, ok := appearences[n]
		if !ok {
			appearences[n] = 1
		} else {
			appearences[n] += 1
		}
	}
	for n, times := range appearences {
		if times%2 == 1 {
			return n
		}
	}
	return 0
}
