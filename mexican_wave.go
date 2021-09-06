package kata

import (
	"unicode"
)

func wave(words string) []string {
	result := []string{}
	for i, c := range []rune(words) {
		if c != ' ' {
			res := []rune(words)
			res[i] = unicode.ToUpper(c)
			result = append(result, string(res))
		}
	}
	return result
}
