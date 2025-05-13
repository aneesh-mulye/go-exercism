package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(f FodderCalculator, n int) (float64, error) {
	total, err := f.FodderAmount(n)
	if err != nil {
		return 0.0, err
	}
	factor, err := f.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	return (float64(total) * factor) / float64(n), nil
}

func ValidateInputAndDivideFood(f FodderCalculator, n int) (float64, error) {
	if n <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(f, n)
}

type InvalidCowsError struct {
	NumCows int
	Message string
}

func (icr *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", icr.NumCows, icr.Message)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{NumCows: numCows,
			Message: "there are no negative cows"}
	}
	if numCows == 0 {
		return &InvalidCowsError{NumCows: numCows,
			Message: "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
