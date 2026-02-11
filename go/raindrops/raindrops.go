package raindrops

import "strconv"

func Convert(number int) string {
	var result string
	var was bool
	if number%3 == 0 {
		result += "Pling"
		was = true
	}
	if number%5 == 0 {
		result += "Plang"
		was = true
	}
	if number%7 == 0 {
		result += "Plong"
		was = true
	}
	if !was {
		result += strconv.Itoa(number)
	}
	return result
}
