package main

import (
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp, 101, 103).part1(100))
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp, 101, 103).part2(100))
}

type solver struct {
	robots []robot
	fs     coor
}

func (s *solver) prep(inp string, fieldW, fieldH int64) *solver {
	for _, line := range strings.Split(inp, "\n") {
		s.robots = append(s.robots, newRobot(line))
	}

	s.fs = coor{fieldH, fieldH}

	return s
}

func (s *solver) part1(wait int64) int {
	sum := 0

	for i := range s.robots {
		s.robots[i].move(wait)
	}

	return sum
}

func (s *solver) part2(wait int64) int {
	sum := 0
	return sum
}

type robot struct {
	pos coor
	vel coor
}

func (r *robot) move(n int64) *robot {

	return r
}

func newRobot(in string) robot {
	p, v, _ := strings.Cut(in, " ")
	_, p, _ = strings.Cut(p, "=")
	_, v, _ = strings.Cut(v, "=")

	return robot{
		pos: parsePair(p),
		vel: parsePair(p),
	}
}

func parsePair(in string) coor {
	x, _ := strconv.ParseInt(strings.Split(in, ",")[0], 10, 64)
	y, _ := strconv.ParseInt(strings.Split(in, ",")[1], 10, 64)
	return coor{x: x, y: y}
}

type coor struct {
	x, y int64
}
