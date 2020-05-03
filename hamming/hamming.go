// Package hamming provides functionality around calculating Hamming distance (link: https://en.wikipedia.org/wiki/Hamming_distance)
package hamming

import (
	"fmt"
)

// Distance returns the Hamming distance between 2 strings or an error if a distance cannot be computed
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("could not compute Hamming distance: %q and %q don't have the same length", a, b)
	}

	dist := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist, nil
}
