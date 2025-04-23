package reverse

func Reverse(input string) string {
	rs := []rune(input)
	for i := 0; i < len(rs)/2; i++ {
		r := rs[i]
		rs[i] = rs[len(rs)-i-1]
		rs[len(rs)-i-1] = r
	}
	return string(rs)
}
