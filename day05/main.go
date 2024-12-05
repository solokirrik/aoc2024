package main

import (
	_ "embed"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Mul:", new(solver).prep(inp).part1())
	slog.Info("Part 2:", "Mul:", new(solver).prep(inp).part2())
}

type solver struct {
	rules   []rule
	manuals [][]int
}

type rule struct {
	x, y int
}

func (s *solver) prep(inp string) *solver {
	parts := strings.Split(inp, "\n\n")
	rawRules := strings.Split(parts[0], "\n")
	for i := range rawRules {
		newRulesStr := strings.Split(rawRules[i], "|")
		x, _ := strconv.Atoi(newRulesStr[0])
		y, _ := strconv.Atoi(newRulesStr[1])
		s.rules = append(s.rules, rule{x: x, y: y})
	}

	rawManuals := strings.Split(parts[1], "\n")
	s.manuals = make([][]int, 0, len(rawManuals))
	for i := range rawManuals {
		newRowStr := strings.Split(rawManuals[i], ",")
		newRow := make([]int, 0, len(newRowStr))
		for j := range newRowStr {
			val, _ := strconv.Atoi(newRowStr[j])
			newRow = append(newRow, val)
		}

		s.manuals = append(s.manuals, newRow)
	}

	return s
}

func (s *solver) part1() int {
	sum := 0
	for _, manual := range s.manuals {
		if s.isValid(manual) {
			sum += manual[len(manual)/2]
		}
	}

	return sum
}

func (s *solver) isValid(manual []int) bool {
	for i := 0; i < len(manual); i++ {
		val := manual[i]
		options := lo.Filter(s.rules, func(ruleVal rule, _ int) bool {
			return ruleVal.x == val
		})

		for _, option := range options {
			if slices.Contains(manual[:i], option.y) {
				return false
			}
		}
	}

	return true
}

func (s *solver) part2() int {
	sum := 0
	for _, manual := range s.manuals {
		if s.isValid(manual) {
			continue
		}

		manual = s.correct(manual)
		midVal := manual[len(manual)/2]
		sum += midVal
	}

	return sum
}

func (s *solver) correct(manual []int) []int {
	slices.SortFunc(manual, func(a, b int) int {
		switch {
		case slices.Contains(s.rules, rule{a, b}):
			return 1
		case slices.Contains(s.rules, rule{b, a}):
			return -1
		default:
			return 0
		}
	})
	return manual
}
