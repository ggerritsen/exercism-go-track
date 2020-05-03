// Package luhn provides functionality around checking Luhn validity (https://en.wikipedia.org/wiki/Luhn_algorithm)
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns whether a string is valid according to Luhn's algorithm
func Valid(s string) bool {
	s2 := strings.ReplaceAll(s, " ", "")
	if len(s2) < 2 {
		return false
	}

	// sanitize input
	var v []int
	for _, l := range s2 {
		i, err := strconv.Atoi(string(l))
		if err != nil {
			return false
		}

		v = append(v, i)
	}

	for i := len(v) - 2; i >= 0; i = i - 2 {
		v[i] = v[i] * 2
		if v[i] > 9 {
			v[i] = v[i] - 9
		}
	}

	sum := 0
	for _, vv := range v {
		sum += vv
	}

	return sum%10 == 0
}
