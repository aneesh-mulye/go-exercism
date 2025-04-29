package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	Hours   int
	Minutes int
}

// Not handling negative b because that never happens in this exercise
func mod(a, b int) int {
	return (b + a%b) % b
}

func div(a, n int) int {
	ret := a / n
	if a <= 0 && mod(a, n) != 0 {
		ret--
	}
	return ret
}

func New(h, m int) Clock {
	return Clock{
		Hours:   mod(h+div(m, 60), 24),
		Minutes: mod(m, 60),
	}
}

func (c Clock) Add(m int) Clock {
	m += c.Minutes
	return Clock{
		Hours:   mod(c.Hours+div(m, 60), 24),
		Minutes: mod(m, 60),
	}
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}
