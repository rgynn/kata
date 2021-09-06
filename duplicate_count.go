package kata

import "strings"

func duplicate_count(s1 string) int {
	dups := map[rune]bool{}
	for _, c := range strings.ToLower(s1) {
		_, ok := dups[c]
		if !ok {
			dups[c] = false
		} else {
			dups[c] = true
		}
	}
	var sum int
	for _, dup := range dups {
		if dup {
			sum++
		}
	}
	return sum
}
