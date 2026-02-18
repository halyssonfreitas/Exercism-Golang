package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(sentence string) Frequency {
	result := make(Frequency)
	reg := regexp.MustCompile(`\w+('\w+)?`)
	for _, word := range reg.FindAllString(strings.ToLower(sentence), -1) {
		result[word]++
	}
	return result
}
