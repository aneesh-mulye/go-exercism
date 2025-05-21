package prime

import "errors"

func isPrime(candidate int, primesSoFar []int) bool {
	for _, primeFactor := range primesSoFar {
		if primeFactor*primeFactor > candidate {
			return true
		}
		if candidate%primeFactor == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n in nth prime must be a positive natural")
	}
	// Special-casing this so the inner loop logic is faster/less convoluted
	if n == 1 {
		return 2, nil
	}
	primesSoFar := []int{2}
	for candidate, primesFound := 3, 1; primesFound < n; candidate += 2 {
		if isPrime(candidate, primesSoFar) {
			primesFound++
			primesSoFar = append(primesSoFar, candidate)
		}
	}
	return primesSoFar[len(primesSoFar)-1], nil
}
