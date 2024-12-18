package main

import (
	_ "embed"
	"log/slog"
	"strconv"
	"strings"

	"github.com/solokirrik/aoc2024/pkg"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp).part1())
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp).part2())
}

type pos struct {
	r, c int
	val  int64
}

type topoMap [][]pos

type solver struct {
	mtx    topoMap
	starts []pos
}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n")
	s.mtx = make([][]pos, len(rows))

	for i := range rows {
		row := strings.Split(rows[i], "")
		s.mtx[i] = make([]pos, len(row))

		for j := range row {
			val, err := strconv.ParseInt(row[j], 10, 64)
			pkg.PanicOnErr(err)
			s.mtx[i][j] = newPos(i, j, val)
			if val == 0 {
				s.starts = append(s.starts, newPos(i, j, val))
			}
		}
	}

	return s
}

func (s *solver) part1() int {
	sum := 0
	tracers := make([]tracer, len(s.starts))

	for i, start := range s.starts {
		sum += tracers[i].addMap(s.mtx).move(start).getScore()
	}

	return sum
}

func (s *solver) part2() int {
	sum := 0
	tracers := make([]tracer, len(s.starts))

	for i, start := range s.starts {
		sum += len(tracers[i].addMap(s.mtx).move2(start, []pos{}).paths)
	}

	return sum
}

type tracer struct {
	mtx     topoMap
	visited map[pos]int
	score   int
	paths   [][]pos
}

func (t *tracer) addMap(mtx [][]pos) *tracer {
	t.mtx = mtx
	t.visited = make(map[pos]int)
	t.score = 0

	return t
}

func (t *tracer) move(loc pos) *tracer {
	t.visited[loc]++

	if loc.val == 9 {
		t.score++
		return t
	}

	opt := t.getOptions(loc)
	for _, o := range opt {
		t.move(o)
	}

	return t
}

func (t *tracer) move2(loc pos, curPath []pos) *tracer {
	curPath = append(curPath, loc)

	if loc.val == 9 {
		pathCopy := make([]pos, len(curPath))
		copy(pathCopy, curPath)
		t.paths = append(t.paths, pathCopy)
		return t
	}

	opt := t.getOptions(loc)
	for _, o := range opt {
		t.move2(o, curPath)
	}

	return t
}

func (t *tracer) getScore() int {
	return t.score
}

func (t *tracer) getOptions(cur pos) []pos {
	out := make([]pos, 0, 4)

	deltas := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for _, d := range deltas {
		option := pos{r: cur.r + d[0], c: cur.c + d[1]}
		if !t.mtx.isInBound(option) {
			continue
		} else {
			option.val = t.mtx[option.r][option.c].val
			if t.deltaAceptable(cur, option) && !t.isVisited(option) {
				out = append(out, option)
			}
		}
	}

	return out
}

func (t *tracer) isVisited(p pos) bool {
	return t.visited[p] > 0
}

func (m topoMap) isInBound(p pos) bool {
	return p.r >= 0 && p.c >= 0 && p.r < len(m) && p.c < len(m[0])
}

func (t *tracer) deltaAceptable(c, p pos) bool {
	return p.val-c.val == 1
}

func newPos(row, col int, val int64) pos {
	return pos{r: row, c: col, val: val}
}
