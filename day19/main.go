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
	s := new(solver).prep(inp)

	start := time.Now()
	got1 := s.part1()
	slog.Info("Part 1", "time", time.Since(start).String(), "Ans", got1)

	start2 := time.Now()
	got2 := s.part2()
	slog.Info("Part 2", "time", time.Since(start2).String(), "Ans", got2)
}

type solver struct {
	patterns []string
	designs  []string
}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n\n")
	s.patterns = strings.Split(rows[0], ", ")
	s.designs = strings.Split(rows[1], "\n")
	return s
}

func filterPatterns(design string, patterns []string) []string {
	out := make([]string, 0, len(patterns))

	for i := range patterns {
		if strings.Contains(design, patterns[i]) {
			out = append(out, patterns[i])
		}
	}

	return out
}

func (s *solver) part1() int {
	res := 0
	cache := make(map[string]int, len(s.patterns))

	for i := range s.designs {
		designPatterns := filterPatterns(s.designs[i], s.patterns)
		if val := countVariations(cache, s.designs[i], designPatterns); val > 0 {
			res++
		}
	}

	return res
}

func (s *solver) part2() int {
	res := 0
	cache := make(map[string]int, len(s.patterns))

	for i := range s.designs {
		designPatterns := filterPatterns(s.designs[i], s.patterns)
		res += countVariations(cache, s.designs[i], designPatterns)
	}

	return res
}

func countVariations(cache map[string]int, design string, patterns []string) int {
	if len(design) == 0 {
		return 1
	}

	if val, ok := cache[design]; ok {
		return val
	}

	opts := make([]string, 0, len(patterns))
	for i := range patterns {
		if strings.HasPrefix(design, patterns[i]) {
			opts = append(opts, patterns[i])
		}
	}

	if len(opts) == 0 {
		return 0
	}

	for i := range opts {
		cache[design] += countVariations(cache, design[len(opts[i]):], patterns)
	}

	return cache[design]
}
