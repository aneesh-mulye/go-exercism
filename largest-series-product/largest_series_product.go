package lsproduct

import (
	"errors"
	"unicode"
)

// The existence of zeroes messes up the sliding window. I'm sure there's some
// way of tracking that, but screw it, no.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span <= 0 {
		return 0, errors.New("span must be positive")
	}
	if span > len(digits) {
		return 0, errors.New("span > length")
	}
	for _, char := range digits {
		if !unicode.IsDigit(char) {
			return 0, errors.New("invalid input; expect only digits!")
		}
	}
	var maxProd int64
	for i := 0; i <= len(digits)-span; i++ {
		var tempProd int64 = 1
		for j := i; j < i+span; j++ {
			tempProd *= int64(digits[j] - '0')
		}
		if tempProd > maxProd {
			maxProd = tempProd
		}
	}
	return maxProd, nil
}
