package kata

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) < 1 {
		return 0
	}
	var largest int
	for i := range numbers {
		var sum int
		for _, n := range numbers[i:] {
			sum += n
			if sum > largest {
				largest = sum
			}
		}
	}
	if largest < 0 {
		return 0
	}
	return largest
}
