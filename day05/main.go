package main

import (
	_ "embed"
	"log/slog"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Mul:", new(solver).prep(inp).part1())
	slog.Info("Part 2:", "Mul:", new(solver).prep(inp).part2())
}

type solver struct{}

func (s *solver) prep(_ string) *solver {
	return s
}

func (s *solver) part1() int {
	return 0
}

func (s *solver) part2() int {
	return 0
}
