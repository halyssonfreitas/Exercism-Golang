package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (foodPerCow float64, err error) {
	fa, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	ff, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fa / float64(cows) * ff, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (foodPerCow float64, err error) {
	if cows < 1 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

// TODO: define the 'ValidateNumberOfCows' function

type InvalidCowsError struct {
	cows    int
	message string
}

func NewInvalidCowsError(cows int, message string) *InvalidCowsError {
	return &InvalidCowsError{
		cows:    cows,
		message: message,
	}
}

func (i *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", i.cows, i.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return NewInvalidCowsError(cows, "there are no negative cows")
	}
	if cows == 0 {
		return NewInvalidCowsError(cows, "no cows don't need food")
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
