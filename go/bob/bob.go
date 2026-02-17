// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	allCapitalize := false
	if remark == strings.ToUpper(remark) {
		allCapitalize = true
	}
	var isQuestion bool
	if len(remark) > 0 {
		isQuestion = remark[len(remark)-1] == '?'
		remark = remark[:len(remark)-1]
	}

	hasLetter := regexp.MustCompile("[a-zA-Z]+").Match([]byte(remark))

	fmt.Println("allCapitalize " + strconv.FormatBool(allCapitalize))
	fmt.Println("hasLetter " + strconv.FormatBool(hasLetter))
	fmt.Println("isQuestion " + strconv.FormatBool(isQuestion))

	switch {
	case isQuestion && allCapitalize && hasLetter:
		return "Calm down, I know what I'm doing!"
	case allCapitalize && hasLetter:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case strings.TrimSpace(remark) == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
