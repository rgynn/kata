package kata

import "unicode"

func alphanumeric(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if !unicode.IsNumber(c) && !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}
