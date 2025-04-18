package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, errors.New("require a positive input")
	}

	var iters int
	for n != 1 {
		iters++
		if 0 == n%2 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}

	return iters, nil
}
