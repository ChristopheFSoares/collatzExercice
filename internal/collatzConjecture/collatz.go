package collatzConjecture

import "errors"

func Collatz(value int) (int, error) {
	if value <= 0 {
		return 0, errors.New("invalid value")
	}

	if value == 1 {
		return 1, nil
	}

	count := 0
	for value != 1 {
		if value%2 == 0 {
			value = value / 2

		} else {
			value = 3*value + 1
		}

		count++
	}

	return count, nil
}
