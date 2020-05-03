// Package reverse provides functionality around reversing strings
package reverse

// Reverse returns a string in reverse
func Reverse(s string) string {
	result := ""
	for _, l := range s {
		result = string(l) + result
	}
	return result
}
