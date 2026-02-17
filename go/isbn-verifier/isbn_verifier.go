package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	cleanIsbn := strings.ReplaceAll(isbn, "-", "")
	if len(cleanIsbn) != 10 {
		return false
	}
	var sumOfMultiples int
	for i := 0; i < 9; i++ {
		integer, err := strconv.Atoi(cleanIsbn[i : i+1])
		if err != nil {
			return false
		}
		sumOfMultiples += integer * (10 - i)
	}
	var last int
	if cleanIsbn[9:10] == string("x") {
		last = 10
	} else {
		last, _ = strconv.Atoi(cleanIsbn[9:10])
	}

	sumOfMultiples += last

	if sumOfMultiples%11 != 0 {
		return false
	}

	return true

}
