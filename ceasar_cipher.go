package kata

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func MovingShift(s string, shift int) []string {
	return split(encode(s, shift), 5)
}

func DemovingShift(arr []string, shift int) string {
	return decode(strings.Join(arr, ""), shift)
}

func encode(s string, shift int) (ret string) {
	for i, c := range s {
		if unicode.IsLetter(c) {
			switch {
			case c >= 'A' && c <= 'Z':
				c += int32((i + shift) % 26)
				if c > 'Z' {
					c -= 26
				}
			case c >= 'a' && c <= 'z':
				c += int32((i + shift) % 26)
				if c > 'z' {
					c -= 26
				}
			}
		}
		ret += fmt.Sprintf("%c", c)
	}
	return
}

func decode(s string, shift int) (ret string) {
	for i, c := range s {
		if unicode.IsLetter(c) {
			switch {
			case c >= 'A' && c <= 'Z':
				c -= int32((i + shift) % 26)
				if c < 'A' {
					c += 26
				}
			case c >= 'a' && c <= 'z':
				c -= int32((i + shift) % 26)
				if c < 'a' {
					c += 26
				}
			}
		}
		ret += fmt.Sprintf("%c", c)
	}
	return
}

func split(s string, p int) (result []string) {
	result = make([]string, p)
	d := int(math.Ceil(float64(len(s)) / float64(p)))
	for i := 0; i < p; i++ {
		x := int64(math.Min(float64(i*d), float64(len(s))))
		y := int64(math.Min(float64((i+1)*d), float64(len(s))))
		result[i] = s[x:y]
	}
	return
}
