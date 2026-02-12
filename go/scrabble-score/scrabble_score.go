package scrabble

import "strings"

func Score(word string) int {
	var count int
	wordLowerCase := strings.ToLower(word)
	for i := 0; i < len(wordLowerCase); i++ {
		switch wordLowerCase[i] {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			count += 1
		case 'd', 'g':
			count += 2
		case 'b', 'c', 'm', 'p':
			count += 3
		case 'f', 'h', 'v', 'w', 'y':
			count += 4
		case 'k':
			count += 5
		case 'j', 'x':
			count += 8
		case 'q', 'z':
			count += 10
		}
	}
	return count
}
