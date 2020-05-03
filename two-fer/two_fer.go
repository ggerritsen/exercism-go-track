// Package twofer provides functionality around sharing between 2 people
package twofer

import "fmt"

// ShareWith returns a string depicting how things will be shared between `me` and `name`
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
