package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	// Lowercase.
	phrase = strings.ToLower(phrase)
	// Split on punctuation.
	re := regexp.MustCompile(`[":?.,!@#$%^&*|\s]+`)
	words := re.Split(phrase, -1)
	// Trim apostraphes or punctuation *from the ends*.
	for i, w := range words {
		words[i] = strings.Trim(w, "'")
	}
	freqs := make(Frequency)
	for _, w := range words {
		freqs[w]++
	}
	delete(freqs, "")
	return freqs
}
