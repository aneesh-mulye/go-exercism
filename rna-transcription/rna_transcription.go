package strand

var dnaToRna map[rune]rune = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'A': 'U',
	'T': 'A',
}

func ToRNA(dna string) string {
	rna := make([]rune, len(dna), len(dna))
	for i, nuc := range dna {
		rna[i] = dnaToRna[nuc]
	}
	return string(rna)
}
