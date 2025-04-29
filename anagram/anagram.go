package anagram

import (
	"strings"
)

func lfreqs(s string) map[rune]int {
	s = strings.ToLower(s)
	ret := make(map[rune]int)
	for _, r := range s {
		ret[r]++
	}
	return ret
}

func containedp(s, other map[rune]int) bool {
	for k, v := range s {
		if v <= other[k] {
			continue
		}
		return false
	}
	return true
}

func anagramp(s, other map[rune]int) bool {
	return containedp(s, other) && containedp(other, s)
}

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	sfreqs := lfreqs(subject)
	ret := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		if subject == strings.ToLower(candidate) {
			continue
		}
		if anagramp(sfreqs, lfreqs(candidate)) {
			ret = append(ret, candidate)
		}
	}
	return ret
}
