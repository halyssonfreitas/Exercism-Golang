package isogram

import "strings"

func IsIsogram(word string) bool {
	isogramMap := make(map[rune]int, 0)

	word = strings.ToLower(word)

	for _, r := range word {
		if _, ok := isogramMap[r]; ok {
			if r != ' ' && r != '-' {
				return false
			}
		}
		isogramMap[r] = 1
	}

	return true
}
