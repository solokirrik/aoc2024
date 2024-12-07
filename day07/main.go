package main

import (
	_ "embed"
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp).part1())
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp).part2())
}

type solver struct {
	equations []equation
}

type equation struct {
	res  uint64
	nums []uint64
}

var signs = []string{"+", "*"}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n")
	s.equations = make([]equation, 0, len(rows))

	for _, row := range rows {
		res, nums, _ := strings.Cut(row, ": ")

		eqRes, err := strconv.ParseUint(res, 10, 64)
		panicIsErr(err)

		eq := equation{
			res:  eqRes,
			nums: make([]uint64, 0, strings.Count(nums, " ")+1),
		}

		for _, item := range strings.Split(nums, " ") {
			val, err := strconv.ParseUint(item, 10, 64)
			panicIsErr(err)
			eq.nums = append(eq.nums, val)
		}

		s.equations = append(s.equations, eq)
	}

	return s
}

func (s *solver) part1() int {
	sum := uint64(0)

	for _, eq := range s.equations {
		if len(eq.nums) == 1 && eq.res == eq.nums[0] {
			sum += eq.nums[0]
			continue
		}

		for _, opt := range getBinCombinations(len(eq.nums) - 1) {
			if eq.res == calculateOption(opt, eq.nums) {
				sum += eq.res
				break
			}
		}
	}

	return int(sum)
}

func (s *solver) part2() int {
	sum := uint64(0)

	for _, eq := range s.equations {
		for _, opt := range getTriCombinations(len(eq.nums) - 1) {
			if eq.res == calculateOption(opt, eq.nums) {
				sum += eq.res
				break
			}
		}
	}

	return int(sum)
}

const (
	PLUS   = 0
	MUL    = 1
	CONCAT = 2
)

func calculateOption(signs []int, nums []uint64) uint64 {
	res := uint64(nums[0])
	for i := 0; i < len(signs); i++ {
		switch signs[i] {
		case CONCAT:
			res = concat(res, nums[i+1])
		case MUL:
			res = res * nums[i+1]
		case PLUS:
			res = res + nums[i+1]
		}
	}

	return res
}

func concat(a, b uint64) uint64 {
	return a*pow(uint64(10), countDigits(b)) + b
}

func countDigits(num uint64) int {
	if num == 0 {
		return 1
	}
	return int(math.Log10(math.Abs(float64(num))) + 1)
}

func getBinCombinations(length int) [][]int {
	totalCombinations := pow(2, length)
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

func getTriCombinations(length int) [][]int {
	totalCombinations := pow(3, length)
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

func pow[T, U int | int64 | uint | uint64](base T, exp U) T {
	result := T(1)
	for exp > 0 {
		result *= base
		exp--
	}

	return result
}

func panicIsErr(err error) {
	if err != nil {
		panic((err))
	}
}
