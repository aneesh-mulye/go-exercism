package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}
	total := 0
	for i, r := range isbn {
		if r == 'X' && i != len(isbn)-1 {
			return false
		}
		if r == 'X' {
			r = '9' + 1
		}
		total += (10 - i) * int(r-'0')
	}
	return total%11 == 0
}
