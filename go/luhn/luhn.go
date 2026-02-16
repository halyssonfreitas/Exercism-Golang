package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.TrimSpace(id)
	oldId := id
	id = ""
	for i := len(oldId) - 1; i >= 0; i-- {
		if string(oldId[i]) != " " {
			id += string(oldId[i])
		}
	}
	length := len((id))
	if length < 2 {
		return false

	}
	var sumOfDoubleThese int
	for i := 0; i < length; i++ {
		integer, err := strconv.Atoi(string(id[i]))
		fmt.Println(integer)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if (i)%2 == 1 {
			if integer*2 > 9 {
				integer = integer*2 - 9
			} else {
				integer = integer * 2
			}
		}
		sumOfDoubleThese += integer
	}
	if sumOfDoubleThese%10 != 0 {
		return false
	}
	return true
}
