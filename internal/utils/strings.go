package utils

import (
	"strings"
	"unicode"
)

func MixCase(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	upper := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			if upper {
				r = unicode.ToUpper(r)
			} else {
				r = unicode.ToLower(r)
			}
			upper = !upper
		}
		b.WriteRune(r)
	}

	return b.String()
}
