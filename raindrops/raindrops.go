// Package raindrops provides functionality around raindrops as strings
package raindrops

import (
	"strconv"
)

// Convert returns a string containing raindrop sounds or if no raindrop generation is possible, the input is returned
func Convert(i int) string {
	result := ""
	if i %3 == 0 {
		result += "Pling"
	}
	if i%5 == 0 {
		result += "Plang"
	}
	if i%7 == 0 {
		result += "Plong"
	}

	// if no factoring is possible, return the input
	if result == "" {
		return strconv.Itoa(i)
	}
	return result
}
