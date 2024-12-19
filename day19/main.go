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
	slog.Info("Starting day18")
	start := time.Now()
	got1 := new(solver).prep(inp).part1()
	slog.Info("Part 1", "time", time.Since(start).String(), "Ans", got1)

	start2 := time.Now()
	got2 := new(solver).prep(inp).part2()
	slog.Info("Part 2", "time", time.Since(start2).String(), "Ans", got2)
}

type solver struct {
	patterns []string
	designs  []string
}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n\n")
	s.patterns = strings.Split(rows[0], ", ")
	s.designs = rows[1:]
	return s
}

func (s *solver) part1() int {
	return 0
}

func (s *solver) part2() int {
	return 0
}
