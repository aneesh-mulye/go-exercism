package rotationalcipher

import (
	"unicode"
	"unicode/utf8"
)

func RotationalCipher(plain string, shiftKey int) string {
	ret := make([]rune, 0, utf8.RuneCountInString(plain))
	for _, r := range plain {
		if !unicode.IsLetter(r) {
			ret = append(ret, r)
			continue
		}
		var base rune
		if unicode.IsUpper(r) {
			base = 'A'
		} else {
			base = 'a'
		}
		ret = append(ret, base+(r-base+rune(shiftKey))%26)
	}
	return string(ret)
}
