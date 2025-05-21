package atbash

import "strings"

func Atbash(s string) string {
	atbash := "zyxwvutsrqponmlkjihgfedcba"
	s = strings.ToLower(s)
	var ret []rune
	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			ret = append(ret, rune(atbash[r-'a']))
		case '0' <= r && r <= '9':
			ret = append(ret, r)
		default:
			continue
		}
	}
	s = string(ret)
	ret = []rune{}
	for i, r := range s {
		ret = append(ret, r)
		if i%5 == 4 {
			ret = append(ret, ' ')
		}
	}
	return strings.TrimRight(string(ret), " ")
}
