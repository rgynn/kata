package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	if n < 1 {
		return []float64{}
	}
	numres := n
	result := []float64{signature[0], signature[1], signature[2]}
	for n > 0 {
		var sum float64
		for x := 3; x > 0; x-- {
			sum += result[len(result)-x]
		}
		result = append(result, sum)
		n--
	}
	return result[:numres]
}
