package chance

import "math/rand/v2"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return (rand.Int() % 19) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	cheap := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(cheap), func(i, j int) {
		iplus := i + 1
		if iplus > 7 {
			iplus -= 7
		}
		jplus := i + 1
		if jplus > 7 {
			jplus -= 7
		}
		aux := cheap[i]
		cheap[i] = cheap[iplus]
		cheap[iplus] = aux
		aux = cheap[j]
		cheap[j] = cheap[jplus]
		cheap[jplus] = aux
	})
	return cheap
}
