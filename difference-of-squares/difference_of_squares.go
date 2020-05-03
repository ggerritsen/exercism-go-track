// Package diffsquares provides functionality around calculating the difference between the sum of the squares
// and the square of the sums of a number
package diffsquares

// SquareOfSum returns the square of the sum of the numbers leading up to `i` (inclusive)
func SquareOfSum(input int) int {
	sum := 0
	for i := 0; i <= input; i++ {
		sum += i
	}
	return sum * sum
}

// SquareOfSum returns the sum of the squares of the numbers leading up to `i` (inclusive)
func SumOfSquares(input int) int {
	sum := 0
	for i := 0; i <= input; i++ {
		sum += i * i
	}
	return sum
}

// Difference returns the difference between the sum of squares and the square of sums for `i`
func Difference(i int) int {
	i1 := SquareOfSum(i)
	i2 := SumOfSquares(i)

	diff := i1 - i2
	if diff < 0 {
		diff = -diff
	}
	return diff
}
