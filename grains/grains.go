// Package grains provides functionality about grains on a chessboard
package grains

import (
	"fmt"
)

// Square provides the amount of wheat grains on a chessboard square
func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, fmt.Errorf("invalid chessboard square: %d", i)
	}

	// unoptimized version:
	//x := math.Pow(float64(2), float64(i-1))
	//return uint64(x), nil

	return 1 << (i - 1), nil
}

// Total provides the total amount of wheat grains on the chessboard
func Total() uint64 {
	//sum := uint64(0)
	//for i := 1; i < 65; i++ {
	//	x, _ := Square(i)
	//	sum += x
	//}
	//return sum

	return 1<<64 - 1
}
