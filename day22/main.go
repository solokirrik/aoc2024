package main

import (
	_ "embed"
	"log/slog"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/solokirrik/aoc2024/pkg"
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
	buyers []uint64
}

func (s *solver) prep(inp string) *solver {
	lines := strings.Split(inp, "\n")
	var err error
	s.buyers = make([]uint64, len(lines))
	for i, line := range lines {
		s.buyers[i], err = strconv.ParseUint(line, 10, 64)
		pkg.PanicOnErr(err)
	}

	return s
}

func (s *solver) part1() uint64 {
	out := uint64(0)

	for i := range s.buyers {
		for j := 0; j < 2000; j++ {
			s.buyers[i] = secretNumber(s.buyers[i])
		}
		out += s.buyers[i]
	}
	return out
}

func (s *solver) part2() int {
	return 0
}

func secretNumber(val uint64) uint64 {
	/*
		Calculate the result of multiplying the secret number by 64.
		Then, mix this result into the secret number.
		Finally, prune the secret number.
	*/
	res := prune(mix(val*64, val))
	/*
		Calculate the result of dividing the secret number by 32.
		Round the result down to the nearest integer.
		Then, mix this result into the secret number.
		Finally, prune the secret number.
	*/
	res = mix(uint64(math.Floor(float64(res)/32.0)), res)
	/*
		Calculate the result of multiplying the secret number by 2048.
		Then, mix this result into the secret number.
		Finally, prune the secret number.
	*/
	res = prune(mix(res*2048, res))
	return res
}

func mix(a, b uint64) uint64 {
	return a ^ b
}

func prune(a uint64) uint64 {
	return a % 16777216
}
