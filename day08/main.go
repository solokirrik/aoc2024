package main

import (
	_ "embed"
	"log/slog"
	"strings"

	"github.com/solokirrik/aoc2024/pkg"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp).part1())
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp).part2())
}

type solver struct {
	mtx      [][]string
	antennas map[string][]pos
}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n")
	s.mtx = make([][]string, 0, len(rows))
	s.antennas = make(map[string][]pos)

	for r := range rows {
		newRow := strings.Split(rows[r], "")
		s.mtx = append(s.mtx, newRow)

		for c, freq := range newRow {
			if freq != "." {
				s.antennas[freq] = append(s.antennas[freq], newPos(r, c))
			}
		}
	}

	return s
}

type pos struct {
	r, c int
}

func newPos(r, c int) pos {
	return pos{r, c}
}

func (s *solver) part1() int {
	sum := make(map[pos]int)

	for _, antennas := range s.antennas {
		for _, pair := range pkg.GeneratePairs(len(antennas)) {
			a := antennas[pair[0]]
			b := antennas[pair[1]]

			if a1 := s.antinode(a, b, 1); s.isOnMap(a1) {
				sum[newPos(a1.r, a1.c)]++
			}
			if a2 := s.antinode(b, a, 1); s.isOnMap(a2) {
				sum[newPos(a2.r, a2.c)]++
			}
		}
	}

	return len(sum)
}

func (s *solver) part2() int {
	sum := make(map[pos]int)

	for _, antennas := range s.antennas {
		for _, pair := range pkg.GeneratePairs(len(antennas)) {
			a := antennas[pair[0]]
			b := antennas[pair[1]]

			//are all also antinodes
			sum[a]++
			sum[b]++

			for _, an := range s.antinodes(a, b) {
				sum[an]++
			}
			for _, an := range s.antinodes(b, a) {
				sum[an]++
			}
		}
	}

	return len(sum)
}

func (s *solver) isOnMap(a pos) bool {
	return a.r >= 0 && a.c >= 0 && a.r <= len(s.mtx[0])-1 && a.c <= len(s.mtx)-1
}

func (s *solver) antinodes(center, b pos) []pos {
	out := []pos{}

	i := 1
	for {
		an := s.antinode(center, b, i)
		if s.isOnMap(an) {
			out = append(out, an)
			i++
		} else {
			break
		}
	}

	return out
}

func (s *solver) antinode(center, b pos, i int) pos {
	dx := pkg.Abs(max(center.c, b.c) - min(center.c, b.c))
	dy := pkg.Abs(max(center.r, b.r) - min(center.r, b.r))

	anX := center.c + i*dx
	anY := center.r + i*dy

	if center.r < b.r {
		anY = center.r - i*dy
	}
	if center.c < b.c {
		anX = center.c - i*dx
	}

	return newPos(anY, anX)
}
