package pangram

import "unicode"

func IsPangram(input string) bool {
	present := make(map[rune]bool)
	for _, r := range input {
		present[unicode.ToLower(r)] = true
	}
	for r := 'a'; r <= 'z'; r++ {
		if !present[r] {
			return false
		}
	}
	return true
}
