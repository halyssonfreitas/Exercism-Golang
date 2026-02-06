package logs

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	// for _, c := range log {
	// 	switch c {
	// 	case '❗':
	// 		return "recommendation"
	// 	case '🔍':
	// 		return "search"
	// 	case '☀':
	// 		return "weather"
	// 	}
	// }
	// return "default"
	for _, runeC := range []rune(log) {
		switch runeC {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	logRune := []rune(log)
	for i, c := range logRune {
		if c == oldRune {
			logRune[i] = newRune
		}
	}
	return string(logRune)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	var lengthString int
	lengthString = len(log)
	lengthString = lengthString
	lengthRunner := utf8.RuneCountInString(log)
	fmt.Println("Len count  : " + strconv.Itoa(lengthString))
	fmt.Println("Rune count : " + strconv.Itoa(lengthRunner))
	if lengthRunner > limit {
		return false
	}
	return true
}
