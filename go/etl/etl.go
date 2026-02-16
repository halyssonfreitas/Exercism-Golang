package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int, 0)

	for key, value := range in {
		for _, s := range value {
			l := strings.ToLower(s)
			result[l] = key
		}
	}

	return result
}
