package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("")
	}
	if n == 1 {
		return 0, nil
	}
	var steps int
	for {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		steps++
		if n == 1 {
			return steps, nil
		}
	}

}
