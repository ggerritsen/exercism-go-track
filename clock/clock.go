// Package clock provides a clock that handles times without dates
package clock

import "fmt"

type Clock interface {
	// String returns a string representation of a Clock
	String() string
	// Add adds `m` minutes to the Clock's time
	Add(m int) Clock
	// Subtract subtracts `m` minutes from the Clock's time
	Subtract(m int) Clock
}

type myClock struct {
	h, m int
}

func (c myClock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c myClock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c myClock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

// New creates a new Clock
func New(h, m int) Clock {
	// correct minutes >= 60
	if m >= 60 {
		h += m / 60
		m = m % 60
	}

	// correct negative minutes
	if m < 0 {
		mins := -m

		minHours := mins / 60
		minMins := mins % 60

		h -= minHours
		if minMins == 0 {
			m = 0
		} else {
			h--
			m = 60 - minMins
		}
	}

	// correct negative hours
	if h < 0 {
		h = 24 + (h%24)
	}

	// correct hours >= 24
	if h >= 24 {
		h = h % 24
	}

	return myClock{h, m}
}
