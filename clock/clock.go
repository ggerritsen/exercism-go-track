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
	m int
}

func (c myClock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}

func (c myClock) Add(m int) Clock {
	return New(c.m/60, c.m%60+m)
}

func (c myClock) Subtract(m int) Clock {
	return c.Add(-m)
}

// New creates a new Clock
func New(h, m int) Clock {
	mins := h*60 + m

	// remove the days
	mins %= 24 * 60
	if mins < 0 {
		mins += 24 * 60
	}
	return myClock{mins}
}
