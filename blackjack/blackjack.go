package blackjack

/* The actually correct way to do it, but whatever, I'll go with the spirit of
 * the exercise.
var cardValues = map[string]int{
	"ace":   11,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"nine":  9,
	"eight": 8,
	"seven": 7,
	"six":   6,
	"five":  5,
	"four":  4,
	"three": 3,
	"two":   2,
}
*/

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handValue := ParseCard(card1) + ParseCard(card2)
	switch {
	case "ace" == card1 && "ace" == card2:
		return "P"
	case 21 == handValue:
		if ParseCard(dealerCard) >= 10 {
			return "S"
		}
		return "W"
	case 17 <= handValue && handValue <= 20:
		return "S"
	case 12 <= handValue && handValue <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
