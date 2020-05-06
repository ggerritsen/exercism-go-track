// Package matrix provides functionality around matrices
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix provides basic matrix functionality
type Matrix interface {
	// Rows returns a copy of the matrix rows in top-to-bottom order
	Rows() [][]int
	// Cols returns a copy of the matrix columns in left-to-right order
	Cols() [][]int
	// Set sets the value of a cell in the matrix, as defined by `row` and `col`
	Set(row, col, val int) bool
}

type myMatrix [][]int

func (m myMatrix) Rows() [][]int {
	cp := newMatrix(len(m), len(m[0]))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			cp[i][j] = m[i][j]
		}
	}

	return cp
}

func (m myMatrix) Cols() [][]int {
	transposition := newMatrix(len(m[0]), len(m))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			transposition[j][i] = m[i][j]
		}
	}

	return transposition
}

func newMatrix(r, c int) myMatrix {
	m := make([][]int, r)
	for i := 0; i < r; i++ {
		m[i] = make([]int, c)
	}
	return m
}

func (m myMatrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row > len(m)-1 || col > len(m[0])-1 {
		return false
	}

	m[row][col] = val
	return true
}

// New creates a new matrix
func New(s string) (Matrix, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("invalid matrix %q", s)
	}

	rows := strings.Split(s, "\n")
	width := len(strings.Split(strings.TrimSpace(rows[0]), " "))

	m := newMatrix(len(rows), width)
	for i, row := range rows {
		columns := strings.Split(strings.TrimSpace(row), " ")
		if len(columns) != width {
			return nil, fmt.Errorf("inbalanced matrix %q", s)
		}

		for j, cell := range columns {
			val, err := strconv.Atoi(strings.TrimSpace(cell))
			if err != nil {
				return nil, err
			}

			m.Set(i, j, val)
		}
	}

	return m, nil
}
