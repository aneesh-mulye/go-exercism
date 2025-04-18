// Provides a function to check whether a given year is a leap year or not.
package leap

// year int → bool: true when the year is leap, false otherwise.
func IsLeapYear(year int) bool {
	if 0 == year%400 {
		return true
	}

	if 0 == year%100 {
		return false
	}

	if 0 == year%4 {
		return true
	}

	return false
}
