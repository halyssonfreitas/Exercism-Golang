package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	quantities := make(map[string]int)
	for _, l := range layers {
		quantities[l] = quantities[l] + 1
	}
	return quantities["noodles"] * 50, float64(quantities["sauce"]) * float64(0.2)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (scale []float64) {
	for _, q := range quantities {
		scale = append(scale, q*float64(portions)/2)
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
