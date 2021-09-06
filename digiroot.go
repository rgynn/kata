package kata

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	var result int
	for n != 0 {
		remainder := n % 10
		result += remainder
		n = n / 10
	}
	return DigitalRoot(result)
}
