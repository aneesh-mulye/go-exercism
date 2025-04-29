package proverb

func Proverb(rhyme []string) []string {
	ret := make([]string, 0, len(rhyme))
	if len(rhyme) == 0 {
		return ret
	}
	for i := 0; i < len(rhyme)-1; i++ {
		ret = append(ret,
			"For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
	}
	ret = append(ret, "And all for the want of a "+rhyme[0]+".")
	return ret
}
