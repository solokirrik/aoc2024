package utils

import "golang.org/x/exp/constraints"

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
