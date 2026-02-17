package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 {
		return 0, errors.New("")
	}
	if number > 64 {
		return 0, errors.New("")
	}
	return uint64(math.Pow(2, float64(number)-1)), nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		t, _ := Square(i)
		total += t
	}
	return total
}
