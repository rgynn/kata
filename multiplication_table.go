package kata

func MultiplicationTable(size int) (result [][]int) {
	for x := 1; x <= size; x++ {
		row := []int{}
		for y := 1; y <= size; y++ {
			row = append(row, x*y)
		}
		result = append(result, row)
	}
	return
}
