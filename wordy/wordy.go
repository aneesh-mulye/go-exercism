package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	// First, strip the question mark
	question = strings.ReplaceAll(question, "?", "")
	// Then, split by space
	q := strings.Split(question, " ")
	if len(q) < 3 {
		return 0, false
	}
	ret, err := strconv.Atoi(q[2])
	if err != nil {
		return 0, false
	}
	ops := q[3:]
	for i := 0; i < len(ops); {
		switch ops[i] {
		case "plus":
			// Check that there is in fact another operand in front
			if i+1 >= len(ops) {
				return 0, false
			}
			other, err := strconv.Atoi(ops[i+1])
			if err != nil {
				return 0, false
			}
			ret += other
			i += 2
		case "minus":
			if i+1 >= len(ops) {
				return 0, false
			}
			other, err := strconv.Atoi(ops[i+1])
			if err != nil {
				return 0, false
			}
			ret -= other
			i += 2
		case "multiplied":
			if i+2 >= len(ops) {
				return 0, false
			}
			if ops[i+1] != "by" {
				return 0, false
			}
			other, err := strconv.Atoi(ops[i+2])
			if err != nil {
				return 0, false
			}
			ret *= other
			i += 3
		case "divided":
			if i+2 >= len(ops) {
				return 0, false
			}
			if ops[i+1] != "by" {
				return 0, false
			}
			other, err := strconv.Atoi(ops[i+2])
			if err != nil {
				return 0, false
			}
			ret /= other
			i += 3
		default:
			return 0, false
		}
	}

	return ret, true
}
