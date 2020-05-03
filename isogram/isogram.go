// Package isogram provides functionality around checking whether a word is an isogram (https://en.wikipedia.org/wiki/Isogram)
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns whether or not a word is an isogram
func IsIsogram(s string) bool {
	seen := map[rune]bool{}
	for _, l := range strings.ToLower(s) {
		if !unicode.IsLetter(l) {
			continue
		}

		if seen[l] {
			return false
		}
		seen[l] = true
	}
	return true
}
