package luhn

import (
	"strings"
)

func Valid(id string) bool {
	newId := strings.ReplaceAll(id, " ", "")

	if len(newId) <= 1 {
		return false
	}

	sum := 0
	parity := (len(newId) - 2) % 2

	for i := 0; i < len(newId); i++ {
		n := int(newId[i]) - '0'

		// check if n is a digit
		if n < 0 || n > 9 {
			return false
		}

		if i%2 == parity {
			n *= 2
		}

		if n > 9 {
			n -= 9
		}

		sum += n
	}

	return sum%10 == 0
}
