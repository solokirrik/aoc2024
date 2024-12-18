package main

import (
	_ "embed"
	"log/slog"
	"strconv"
	"strings"

	"github.com/solokirrik/aoc2024/pkg"
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

		for _, opt := range pkg.GetBinCombinations(len(eq.nums) - 1) {
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
		for _, opt := range pkg.GetTriCombinations(len(eq.nums) - 1) {
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
	return a*pkg.Pow(uint64(10), pkg.CountDigits(b)) + b
}

func panicIsErr(err error) {
	if err != nil {
		panic((err))
	}
}
