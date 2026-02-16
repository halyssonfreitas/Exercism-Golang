package darts

import "math"

func Score(x, y float64) int {

	hypotenuse := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	switch {
	case hypotenuse > 10:
		return 0
	case hypotenuse > 5:
		return 1
	case hypotenuse > 1:
		return 5
	default:
		return 10
	}
}
