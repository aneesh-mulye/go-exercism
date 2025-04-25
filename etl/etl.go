package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	ret := map[string]int{}
	for lval, letters := range in {
		for _, letter := range letters {
			ret[strings.ToLower(letter)] = lval
		}
	}
	return ret
}
