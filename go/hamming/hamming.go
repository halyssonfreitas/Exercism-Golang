package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Different lengths should not work")
	}
	var count int
	for i, elem := range a {
		if rune(b[i]) != elem {
			count++
		}
	}
	return count, nil
}
