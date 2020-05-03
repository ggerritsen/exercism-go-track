// Package acronym provides acronym functionality
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns an acronym for `s`
func Abbreviate(s string) string {
	// make sure all hyphenated words count as 2 words
	s2 := strings.ReplaceAll(s, "-", " ")

	result := ""
	for _, w := range strings.Fields(s2) {

		// collect first letter of each word and uppercase it
		for _, l := range w {
			if unicode.IsLetter(l) {
				result = result + string(l)
				break
			}
		}
	}

	return strings.ToUpper(result)
}
