package pangram

import "strings"

func IsPangram(input string) bool {

	input = strings.ToLower(input)

	alphabet := map[rune]bool{}
	letter := 'a'
	for {
		if letter > 'z' {
			break
		}
		alphabet[letter] = false
		letter++
	}
	for _, l := range input {
		if _, ok := alphabet[l]; ok {
			alphabet[l] = true
		}

	}
	for _, v := range alphabet {
		if !v {
			return false
		}
	}
	return true
}
