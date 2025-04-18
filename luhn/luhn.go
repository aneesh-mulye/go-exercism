package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	// Strip spaces before processing
	id = strings.ReplaceAll(id, " ", "")

	// Strings of length 1 or less are not valid.
	if len(id) <= 1 {
		return false
	}

	// Make sure it's all digits
	for _, letter := range id {
		if !unicode.IsDigit(letter) {
			return false
		}
	}

	// Figure out which alternating letters to double
	altMod := len(id) % 2
	var luhnTotal int
	for i, letter := range id {
		if altMod == i%2 {
			d, _ := strconv.Atoi(string(letter))
			d *= 2
			if d > 9 {
				d -= 9
			}
			luhnTotal += d
		} else {
			d, _ := strconv.Atoi(string(letter))
			luhnTotal += d
		}
	}

	return 0 == (luhnTotal % 10)
}
