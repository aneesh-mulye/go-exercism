package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("Invalid Roman numeral.")
	}
	ret := ""
	rem := input

	// The 1000s case
	th := rem / 1000
	if th != 0 {
		rem %= 1000
		ret += strings.Repeat("M", th)
	}

	// The 900 case
	if rem >= 900 {
		rem -= 900
		ret += "CM"
	}

	// The 500 case
	if rem >= 500 {
		rem -= 500
		ret += "D"
	}

	// The 400s case
	if rem >= 400 {
		rem -= 400
		ret += "CD"
	}

	// The 100s case
	hu := rem / 100
	if hu != 0 {
		rem %= 100
		ret += strings.Repeat("C", hu)
	}

	// The 90 case
	if rem >= 90 {
		rem -= 90
		ret += "XC"
	}

	// The 50 case
	if rem >= 50 {
		rem -= 50
		ret += "L"
	}

	// The 40s case
	if rem >= 40 {
		rem -= 40
		ret += "XL"
	}

	// The 10s case
	te := rem / 10
	if te != 0 {
		rem %= 10
		ret += strings.Repeat("X", te)
	}

	// The 9 case
	if rem >= 9 {
		rem -= 9
		ret += "IX"
	}

	// The 5 case
	if rem >= 5 {
		rem -= 5
		ret += "V"
	}

	// The 4 case
	if rem >= 4 {
		rem -= 4
		ret += "IV"
	}

	// The 1s case
	on := rem
	if on > 0 {
		rem -= rem
		ret += strings.Repeat("I", on)
	}

	return ret, nil
}
