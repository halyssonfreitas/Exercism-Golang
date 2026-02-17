package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	// 	| M    | D   | C   | L   | X   | V   | I   |
	// 	| ---- | --- | --- | --- | --- | --- | --- |
	// 	| 1000 | 500 | 100 | 50  | 10  | 5   | 1   |
	if input > 3999 {
		return "", errors.New("The max number is 3999")
	}
	if input < 1 {
		return "", errors.New("The min number is 1")
	}

	thousands := input / 1000
	rest := input % 1000
	hundreds := rest / 100
	rest = rest % 100
	dozens := rest / 10
	units := rest % 10

	var result string

	for i := 0; i < thousands; i++ {
		result += string("M")
	}

	switch {
	case hundreds == 9:
		result += "CM"
	case hundreds >= 5:
		result += "D"
		repetition := hundreds - 5
		for i := 0; i < repetition; i++ {
			result += "C"
		}
	case hundreds == 4:
		result += "CD"
	default:
		for i := 0; i < hundreds; i++ {
			result += "C"
		}
	}

	switch {
	case dozens == 9:
		result += "XC"
	case dozens >= 5:
		result += "L"
		repetition := dozens - 5
		for i := 0; i < repetition; i++ {
			result += "X"
		}
	case dozens == 4:
		result += "XL"
	default:
		for i := 0; i < dozens; i++ {
			result += "X"
		}
	}

	switch {
	case units == 9:
		result += "IX"
	case units >= 5:
		result += "V"
		repetition := units - 5
		for i := 0; i < repetition; i++ {
			result += "I"
		}
	case units == 4:
		result += "IV"
	default:
		for i := 0; i < units; i++ {
			result += "I"
		}
	}

	return result, nil
}
