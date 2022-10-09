package thefarm

import (
	"errors"
	"strconv"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
var SillyNephewError error

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, error := weightFodder.FodderAmount()

	switch {
	case fodder < 0 && (error == ErrScaleMalfunction || error == nil):
		error = errors.New("negative fodder")
	case cows < 0:
		error = errors.New("silly nephew, there cannot be " + strconv.Itoa(cows) + " cows")
	case cows == 0:
		error = errors.New("division by zero")
	case error == ErrScaleMalfunction && fodder > 0:
		fodder *= 2
	}

	if error != nil && error != ErrScaleMalfunction {
		return 0, error
	}

	return fodder / float64(cows), nil
}
