package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	lowerSubject := strings.ToLower(subject)
	result := []string{}

	reference := map[rune]int{}

	for _, r := range lowerSubject {
		reference[r] += 1
	}

	for _, candidate := range candidates {
		if len(lowerSubject) != len(candidate) {
			continue
		}
		if lowerSubject == strings.ToLower(candidate) {
			continue
		}
		candidateMap := map[rune]int{}
		for _, r := range strings.ToLower(candidate) {
			candidateMap[r] += 1
		}
		if len(candidateMap) != len(reference) {
			continue
		}
		isAnagram := true
		for rkey, rvalue := range reference {
			cvalues, cok := candidateMap[rkey]
			if !cok || cvalues != rvalue {
				isAnagram = false
				break
			}
		}
		if !isAnagram {
			continue
		}
		result = append(result, candidate)
	}

	return result
}
