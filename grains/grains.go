package grains

import "errors"

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("Invalid chess square number")
	}
	return uint64(1) << (number - 1), nil
}

func Total() uint64 {
	return ^uint64(0)
}
