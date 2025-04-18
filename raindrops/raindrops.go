package raindrops

import "fmt"

func Convert(number int) string {
	var res string
	if 0 == number%3 {
		res += "Pling"
	}

	if 0 == number%5 {
		res += "Plang"
	}

	if 0 == number%7 {
		res += "Plong"
	}

	if "" == res {
		res = fmt.Sprint(number)
	}

	return res
}
