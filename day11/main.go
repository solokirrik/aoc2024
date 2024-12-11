package main

import (
	_ "embed"
	"log/slog"
	"strconv"
	"strings"

	"github.com/solokirrik/aoc2024/utils"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp).part1(25))
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp).part2(75))
}

type solver struct {
	stones []int64
	cache  map[key]int
}

type key struct {
	level int
	val   int64
}

func (s *solver) prep(inp string) *solver {
	vals := strings.Split(inp, " ")
	s.stones = make([]int64, 0, len(vals))

	for _, sval := range vals {
		val, err := strconv.ParseInt(sval, 10, 64)
		utils.PanicOnErr(err)
		s.stones = append(s.stones, val)
	}

	s.cache = make(map[key]int)

	return s
}

func (s *solver) part1(blinkN int) int {
	for i := 0; i < blinkN; i++ {
		for st, stone := range s.stones {
			opts := applyRule(stone)
			s.stones[st] = opts[0]
			if len(opts) == 1 {
				continue
			}
			s.stones = append(s.stones, opts[1])
		}
	}

	return len(s.stones)
}

func (s *solver) part2(blinkN int) int {
	sum := 0

	for i := range s.stones {
		sum += s.apply(s.stones[i], blinkN)
	}

	return sum
}

func (s *solver) apply(val int64, lvl int) int {
	curKey := key{lvl, val}

	cacheVal, ok := s.cache[curKey]
	if ok {
		return cacheVal
	}

	opts := applyRule(val)
	if lvl == 1 {
		return len(opts)
	}

	if len(opts) == 1 {
		got := s.apply(opts[0], lvl-1)
		s.cache[curKey] = got
		return got
	}

	got := s.apply(opts[0], lvl-1) + s.apply(opts[1], lvl-1)
	s.cache[curKey] = got

	return got
}

func applyRule(n int64) []int64 {
	digits := utils.CountDigits(n)

	switch {
	case n == 0:
		return []int64{1}
	case digits%2 == 0:
		halfTens := utils.Pow(int64(10), digits/2)
		n1 := n / halfTens
		return []int64{n1, n - n1*halfTens}
	default:
		return []int64{n * 2024}
	}
}
