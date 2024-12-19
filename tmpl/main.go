package main

import (
	_ "embed"
	"log/slog"
	"strings"
	"time"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Starting day")

	start := time.Now()
	got1 := new(solver).prep(inp).part1()
	slog.Info("Part 1", "time", time.Since(start).String(), "Ans", got1)

	start2 := time.Now()
	got2 := new(solver).prep(inp).part2()
	slog.Info("Part 2", "time", time.Since(start2).String(), "Ans", got2)
}

type solver struct {
}

func (s *solver) prep(inp string) *solver {
	_ = strings.Split(inp, "\n")

	return s
}

func (s *solver) part1() int {
	return 0
}

func (s *solver) part2() int {
	return 0
}
