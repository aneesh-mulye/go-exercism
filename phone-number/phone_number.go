package phonenumber

import (
	"errors"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	ret := make([]rune, 0)
	for _, ch := range phoneNumber {
		if unicode.IsDigit(ch) {
			ret = append(ret, ch)
		}
	}
	if len(ret) < 10 || 11 < len(ret) {
		return "", errors.New("too many numbers")
	}
	if len(ret) == 11 {
		if ret[0] != '1' {
			return "", errors.New("invalid country code")
		}
		ret = ret[1:]
	}
	if ret[0] < '2' || ret[3] < '2' {
		return "", errors.New("invalid starting digit in area code(s)")
	}
	return string(ret), nil
}

func AreaCode(phoneNumber string) (string, error) {
	numString, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return numString[:3], nil
}

func Format(phoneNumber string) (string, error) {
	numString, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + numString[:3] + ") " + numString[3:6] + "-" + numString[6:], nil
}
