package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number > 0 && number <= 64 {
		return uint64(math.Pow(2, float64(number)-1)), nil
	}

	return 0, errors.New("number invalid")
}

func Total() uint64 {
	sum := uint64(0)

	for i := 1; i <= 64; i++ {
		sum += uint64(math.Pow(2, float64(i)-1))
	}

	return sum
}
