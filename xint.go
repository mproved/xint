package xint

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer](value T) T {
	if value < 0 {
		return -value
	}

	return value
}

// https://en.wikipedia.org/wiki/Exponentiation_by_squaring
func Pow[T constraints.Integer, U constraints.Unsigned](base T, exp U) (result T) {
	result = 1

	for {
		if exp&1 != 0 {
			result *= base
		}

		exp >>= 1

		if exp == 0 {
			break
		}

		base *= base
	}

	return
}
