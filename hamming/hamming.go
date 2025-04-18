package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0,
			errors.New("Hamming distance only defined for equal length sequences")
	}

	var diff int
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
