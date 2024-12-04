package main

import (
	_ "embed"
	"log/slog"
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
	mtx [][]rune
}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n")

	s.mtx = [][]rune{}
	for r := range rows {
		s.mtx = append(s.mtx, lo.Map(strings.Split(rows[r], ""), func(val string, _ int) rune {
			return rune(val[0])
		}))
	}

	return s
}

func (s *solver) part1() int {
	counter := 0
	for i := range len(s.mtx) {
		for j := range len(s.mtx[i]) {
			if s.mtx[i][j] == 'X' {
				counter += s.countXMAS(i, j)
			}
		}
	}

	return counter
}

func (s *solver) part2() int {
	counter := 0

	for i := range len(s.mtx) {
		for j := range len(s.mtx[i]) {
			if s.mtx[i][j] == 'A' {
				if s.countMAS(i, j, SMMS) ||
					s.countMAS(i, j, MSSM) ||
					s.countMAS(i, j, MMSS) ||
					s.countMAS(i, j, SSMM) {
					counter++
				}
			}
		}
	}

	return counter
}

type masPosSolver struct {
	target map[[2]int]rune
}

func (s *solver) countMAS(r, c int, mp masPosSolver) bool {
	for _, point := range masBounds {
		cr, cc := r+point[0], c+point[1]
		if !s.isInBound(cr, cc) {
			return false
		}

		if s.mtx[cr][cc] != mp.target[point] {
			return false
		}
	}

	return true
}

func (s *solver) countXMAS(r, c int) int {
	found := 0

	for _, dir := range xmasPoints {
		for i := 0; i < len(xmas); i++ {
			cr, cc := r+i*dir[0], c+i*dir[1]

			inc := s.countXMASInDir(i, cr, cc)
			if inc == -1 {
				break
			}

			found += inc
		}
	}

	return found
}

func (s *solver) isInBound(r, c int) bool {
	return r >= 0 && c >= 0 && r < len(s.mtx) && c < len(s.mtx[0])
}

func (s *solver) countXMASInDir(i, cr, cc int) int {
	if !s.isInBound(cr, cc) || s.mtx[cr][cc] != xmas[i] {
		return -1
	}

	if s.mtx[cr][cc] != xmas[i] {
		return -1
	}

	if i == len(xmas)-1 && s.mtx[cr][cc] == xmas[len(xmas)-1] {
		return 1
	}

	return 0
}

// S.M
// .A.
// S.M
var SMMS = masPosSolver{
	target: map[[2]int]rune{
		{-1, -1}: 'S',
		{-1, 1}:  'M',
		{1, 1}:   'M',
		{1, -1}:  'S',
	},
}

// M.S
// .A.
// M.S
var MSSM = masPosSolver{
	target: map[[2]int]rune{
		{-1, -1}: 'M',
		{-1, 1}:  'S',
		{1, 1}:   'S',
		{1, -1}:  'M',
	},
}

// M.M
// .A.
// S.S
var MMSS = masPosSolver{
	target: map[[2]int]rune{
		{-1, -1}: 'M',
		{-1, 1}:  'M',
		{1, 1}:   'S',
		{1, -1}:  'S',
	},
}

// S.S
// .A.
// M.M
var SSMM = masPosSolver{
	target: map[[2]int]rune{
		{-1, -1}: 'S',
		{-1, 1}:  'S',
		{1, 1}:   'M',
		{1, -1}:  'M',
	},
}

var (
	xmas       = []rune{'X', 'M', 'A', 'S'}
	xmasPoints = [][2]int{
		{1, 0},   // down
		{0, 1},   // right
		{-1, 0},  // up
		{0, -1},  // left
		{1, 1},   // down-right
		{-1, 1},  // up-right
		{-1, -1}, // up-left
		{1, -1},  // down-left
	}
	masBounds = [][2]int{
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, 1},   // down-right
		{1, -1},  // down-left
	}
)
