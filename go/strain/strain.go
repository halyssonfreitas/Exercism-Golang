package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](a []T, f func(T) bool) []T {
	var result []T
	for _, v := range a {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Discard[T any](a []T, f func(T) bool) []T {
	var result []T
	for _, v := range a {
		if !f(v) {
			result = append(result, v)
		}
	}
	return result
}
