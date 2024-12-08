package utils

import "github.com/samber/lo"

func GeneratePairs(n int) [][2]int {
	pairs := make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, [2]int{i, j})
		}
	}
	return pairs
}

func GetBinCombinations(length int) [][]int {
	totalCombinations := Pow(2, length)
	combinations := make([][]int, 0, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]int, length)
		for j := 0; j < length; j++ {
			combination[length-j-1] = (i >> j) & 1
		}
		combinations = append(combinations, lo.Reverse(combination))
	}

	return combinations
}

func GetTriCombinations(length int) [][]int {
	totalCombinations := Pow(3, length)
	combinations := make([][]int, 0, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]int, length)
		num := i
		for j := length - 1; j >= 0; j-- {
			combination[j] = num % 3
			num /= 3
		}

		combinations = append(combinations, combination)
	}

	return combinations
}
