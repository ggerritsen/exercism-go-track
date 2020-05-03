// Package bob provides functionality around Bob's utterances
package bob

import "strings"

// Hey returns a reply from Bob, based on the remark said to him
func Hey(remark string) string {
	r := strings.TrimSpace(remark)

	if r == "" {
		return "Fine. Be that way!"
	}

	lastChar := r[len(r)-1:]

	// yelling, making sure there's at least 1 letter in remark
	if r == strings.ToUpper(r) && r != strings.ToLower(r) {
		if lastChar == "?" {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if lastChar == "?" {
		return "Sure."
	}

	return "Whatever."
}
