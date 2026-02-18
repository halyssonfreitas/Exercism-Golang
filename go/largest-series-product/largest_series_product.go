package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span < 1 {
		return 0, errors.New("Span need to be natural number")
	}

	if span > len(digits) {
		return 0, errors.New("Span is larger than digits")
	}

	products := make([]int64, len(digits))
	for i := 0; i < len(digits)-span+1; i++ {
		products[i] = 1
		for j := i; j < i+span; j++ {
			digit, err := strconv.Atoi(digits[j : j+1])
			if err != nil {
				return 0, errors.New("There are digits that is NaN.")
			}
			products[i] *= int64(digit)
		}
	}

	maxVal := products[0]
	for i := 1; i < len(products); i++ {
		if products[i] > maxVal {
			maxVal = products[i]
		}
	}
	return maxVal, nil

}
