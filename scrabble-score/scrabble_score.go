package scrabble

import "unicode"

var letterScores map[rune]int

func init() {
	scoreSets := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}

	letterScores = make(map[rune]int)

	for score, letters := range scoreSets {
		for _, letter := range letters {
			letterScores[letter] = score
			letterScores[unicode.ToLower(letter)] = score
		}
	}
}

func Score(word string) int {
	var score int
	for _, letter := range word {
		score += letterScores[letter]
	}

	return score
}
