package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var result int

	result += 10 * colorCode(colors[0])
	result += colorCode(colors[1])

	return result
}

// ColorCode returns the resistance value of the given color.
func colorCode(color string) int {
	switch color {
	case "black":
		return 0
	case "brown":
		return 1
	case "red":
		return 2
	case "orange":
		return 3
	case "yellow":
		return 4
	case "green":
		return 5
	case "blue":
		return 6
	case "violet":
		return 7
	case "grey":
		return 8
	case "white":
		return 9
	default:
		return 0
	}
}
