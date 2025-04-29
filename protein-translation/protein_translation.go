package protein

import "errors"

var ErrStop = errors.New("STOP")
var ErrInvalidBase = errors.New("Invalid base")

var codonMap map[string]string

func init() {
	codonMap = make(map[string]string)
	protMap := map[string][]string{
		"Methionine":    {"AUG"},
		"Phenylalanine": {"UUU", "UUC"},
		"Leucine":       {"UUA", "UUG"},
		"Serine":        {"UCU", "UCC", "UCA", "UCG"},
		"Tyrosine":      {"UAU", "UAC"},
		"Cysteine":      {"UGU", "UGC"},
		"Tryptophan":    {"UGG"},
		"STOP":          {"UAA", "UAG", "UGA"},
	}
	for protein, codons := range protMap {
		for _, codon := range codons {
			codonMap[codon] = protein
		}
	}
}

func FromRNA(rna string) ([]string, error) {
	var ret []string
	for len(rna) > 0 {
		codon := rna[0:3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			return ret, nil
		}
		if err == ErrInvalidBase {
			return nil, ErrInvalidBase
		}
		ret = append(ret, protein)
		rna = rna[3:]
	}
	return ret, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := codonMap[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}
