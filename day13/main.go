package main

import (
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed inp
var inp string

const (
	aCost         = uint64(3)
	bCost         = uint64(1)
	p2PrizeOffset = uint64(10000000000000)
)

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp, 0).part1())
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp, p2PrizeOffset).part2())
}

type solver struct {
	machines []machine
}

func (s *solver) prep(inp string, offset uint64) *solver {
	rawMachines := strings.Split(inp, "\n\n")
	s.machines = make([]machine, 0, len(rawMachines))

	for i := range rawMachines {
		s.machines = append(s.machines, parseMachine(rawMachines[i], offset))
	}

	return s
}

func (s *solver) part1() uint64 {
	out := uint64(0)

	for i := range s.machines {
		an, bk := findSolution(s.machines[i].a, s.machines[i].b, s.machines[i].prize)
		out += (aCost*an + bCost*bk)
	}

	return out
}

func (s *solver) part2() uint64 {
	return s.part1()
}

func findSolution(a, b, p button) (an, bk uint64) {
	n, k := solve(a, b, p)

	if k < 0 || float64(uint64(k)) != k {
		return 0, 0
	}
	if n < 0 || float64(uint64(n)) != n {
		return 0, 0
	}

	return uint64(n), uint64(k)
}

type machine struct {
	a, b  button
	prize button
}

type point struct {
	x, y uint64
}

type button struct {
	x, y uint64
}

func solve(a, b, p button) (an, bk float64) {
	ax, ay := float64(a.x), float64(a.y)
	bx, by := float64(b.x), float64(b.y)
	px, py := float64(p.x), float64(p.y)

	bk = (ax*py - ay*px) / (ax*by - ay*bx)
	an = (px - bx*bk) / ax

	return an, bk
}

func parseMachine(inp string, offset uint64) machine {
	m := machine{}
	rows := strings.Split(inp, "\n")

	m.a = parseButton(rows[0])
	m.b = parseButton(rows[1])
	_, xyPrize, _ := strings.Cut(rows[2], ": ")

	xyParts := strings.Split(xyPrize, ", ")
	m.prize.x, _ = strconv.ParseUint(strings.Split(xyParts[0], "=")[1], 10, 64)
	m.prize.y, _ = strconv.ParseUint(strings.Split(xyParts[1], "=")[1], 10, 64)
	m.prize.x += offset
	m.prize.y += offset

	return m
}

func parseButton(in string) button {
	b := button{}
	_, xy, _ := strings.Cut(in, ": ")
	xyParts := strings.Split(xy, ", ")
	b.x, _ = strconv.ParseUint(strings.Split(xyParts[0], "+")[1], 10, 64)
	b.y, _ = strconv.ParseUint(strings.Split(xyParts[1], "+")[1], 10, 64)
	return b
}
