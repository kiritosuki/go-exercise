package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, cnt int) (float64, error) {
	amount, err := fodderCalculator.FodderAmount(cnt)
	if err != nil {
		return 0, err
	}
	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount * factor / float64(cnt), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cnt int) (float64, error) {
	if cnt <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodderCalculator, cnt)
}

type InvalidCowsError struct {
	count   int
	message string
}

func (invalidCousError *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", invalidCousError.count, invalidCousError.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(count int) error {
	if count < 0 {
		return &InvalidCowsError{
			count:   count,
			message: "there are no negative cows",
		}
	} else if count == 0 {
		return &InvalidCowsError{
			count:   count,
			message: "no cows don't need food",
		}
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
