package isogram

import "strings"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	ignore := map[rune]bool{
		' ': true,
		'-': true,
	}
	for _, letter := range strings.ToLower(word) {
		if ignore[letter] {
			continue
		}
		if seen[letter] {
			return false
		}
		seen[letter] = true
	}
	return true
}
