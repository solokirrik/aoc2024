package main

import (
	_ "embed"
	"log/slog"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Mul:", part1(inp))
	slog.Info("Part 2:", "Mul:", part2(inp))
}

func part1(_ string) int {
	return 0
}

func part2(_ string) int {
	return 0
}
