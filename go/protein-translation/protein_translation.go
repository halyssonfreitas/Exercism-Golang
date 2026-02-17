package protein

import "errors"

var (
	ErrStop        = errors.New("stop execution")
	ErrInvalidBase = errors.New("invalid base provided")
)

func FromRNA(rna string) ([]string, error) {
	result := []string{}
	for i := 0; i < len(rna); i += 3 {
		aminoAcid, err := FromCodon(rna[i : i+3])
		if err != nil && err == ErrStop {
			return result, nil
		}
		if err != nil && err == ErrInvalidBase {
			return result, err
		}
		result = append(result, aminoAcid)
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	// | Codon         Amino Acid    |
	// | ------------------ | ------------- |
	// | AUG                | Methionine    |
	// | UUU, UUC           | Phenylalanine |
	// | UUA, UUG           | Leucine       |
	// | UCU, UCC, UCA, UCG | Serine        |
	// | UAU, UAC           | Tyrosine      |
	// | UGU, UGC           | Cysteine      |
	// | UGG                | Tryptophan    |
	// | UAA, UAG, UGA      | STOP          |
	codonsToAminoAcid := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
	}

	codonsToStop := map[string]string{
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	if _, ok := codonsToStop[codon]; ok {
		return "", ErrStop
	}

	aminoAcid, ok := codonsToAminoAcid[codon]
	if !ok {
		return "", ErrInvalidBase
	}

	return aminoAcid, nil

}
