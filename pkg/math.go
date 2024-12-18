package pkg

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Signed | constraints.Float](v T) T {
	if v < 0 {
		return -v
	}

	return v
}

func Pow[T, U int | int64 | uint | uint64](base T, exp U) T {
	result := T(1)
	for exp > 0 {
		result *= base
		exp--
	}

	return result
}

func CountDigits[T int | int64 | uint | uint64](num T) T {
	if num == 0 {
		return 1
	}
	return T(math.Log10(math.Abs(float64(num))) + 1)
}
