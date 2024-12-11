package main

import (
	_ "embed"
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/solokirrik/aoc2024/utils"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp).part1(25))
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp).part2(25))
}

type solver struct {
	stones *Stone
}

type Stone struct {
	val   int64
	right *Stone
	left  *Stone
}

func (s *Stone) add(ns *Stone) *Stone {
	if s.right == nil {
		s.right = ns
		ns.left = s
		return s
	}

	wasRight := *s.right
	s.right = ns
	ns.left = s
	wasRight.left = ns
	ns.right = &wasRight
	return s
}

func (s *Stone) set(val int64) *Stone {
	s.val = val
	return s
}

func (s *solver) prep(inp string) *solver {
	vals := strings.Split(inp, " ")

	for _, sval := range vals {
		val, err := strconv.ParseInt(sval, 10, 64)
		utils.PanicOnErr(err)
		newStone := Stone{val: val}
		if s.stones == nil {
			s.stones = &newStone
			continue
		}
		curStone := s.stones
		for curStone.right != nil {
			curStone = curStone.right
		}
		curStone.add(&newStone)
	}

	return s
}

func (s *solver) part1(blinkN int) int {
	for i := 0; i < blinkN; i++ {
		curStone := s.stones
		for curStone != nil {
			muts := applyRule(curStone.val)
			curStone.val = muts[0]
			if len(muts) == 1 {
				curStone = curStone.right
				continue
			}
			oldNextStone := curStone.right
			newStone := &Stone{val: muts[1]}
			curStone.add(newStone)
			curStone = oldNextStone
		}
	}

	stones := 1
	curStone := s.stones

	for curStone != nil {
		curStone = curStone.right
		stones++
	}

	return stones
}

func (s *solver) part2(blinkN int) int {
	return 0
}

func applyRule(n int64) []int64 {
	switch {
	case n == 0:
		return []int64{1}
	case countDigits(n)%2 == 0:
		digits := countDigits(n)
		halfTens := utils.Pow(int64(10), digits/2)
		n1 := n / halfTens
		return []int64{n1, n - n1*halfTens}
	default:
		return []int64{n * 2024}
	}
}

func countDigits[T int | int64 | uint | uint64](num T) T {
	if num == 0 {
		return 1
	}
	return T(math.Log10(math.Abs(float64(num))) + 1)
}
