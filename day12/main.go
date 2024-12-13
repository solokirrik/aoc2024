package main

import (
	"bytes"
	_ "embed"
	"log/slog"

	"github.com/samber/lo"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp).part1())
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp).part2())
}

type solver struct {
	rows    [][]byte
	visited map[point]int
	shapes  []shape
}

func newShape(id int, start point, mtx [][]byte) shape {
	return shape{
		id:       id,
		queuePos: []point{start},
		mtx:      mtx,
		intern:   make(map[point]int),
		border:   make(map[point]int),
		outer:    make(map[point]int),
		visited:  make(map[dirPoint]int),
	}
}

func (s *solver) prep(inp string) *solver {
	rawRows := bytes.Split([]byte(inp), []byte("\n"))
	emptyFirstLine := make([]byte, len(rawRows[0])+2)
	emptyLastLine := make([]byte, len(rawRows[0])+2)

	for i := range rawRows {
		rawRows[i] = append(append([]byte{0}, rawRows[i]...), 0)
	}

	s.rows = [][]byte{emptyFirstLine}
	s.rows = append(s.rows, rawRows...)
	s.rows = append(s.rows, emptyLastLine)

	s.visited = make(map[point]int)

	return s
}

func (s *solver) addShape(sh shape) {
	s.shapes = append(s.shapes, sh)
	for k, v := range sh.intern {
		s.visited[k] += v
	}
}

func (s *solver) part1() int {
	sum := 0
	for i := range s.rows {
		for j := range s.rows {
			newPoint := point{i, j, s.rows[i][j]}
			if s.rows[i][j] == 0 || s.visited[newPoint] > 0 {
				continue
			}

			sh := newShape(len(s.shapes), newPoint, s.rows)
			sh.aggShape()
			s.addShape(sh)

			sum += len(sh.intern) * sh.getPerimeter()
		}
	}

	return sum
}

func (s *solver) part2() int {
	sum := 0

	for i := range s.rows {
		for j := range s.rows {
			newPoint := point{i, j, s.rows[i][j]}
			if s.rows[i][j] == 0 || s.visited[newPoint] > 0 {
				continue
			}

			sh := newShape(len(s.shapes), newPoint, s.rows)
			sh.aggShape()
			s.addShape(sh)

			sum += len(sh.intern) * sh.countSides()
		}
	}

	return sum
}

func (sh *shape) getPerimeter() int {
	perimeter := 0
	for _, v := range sh.outer {
		perimeter += v
	}
	return perimeter
}

func (sh *shape) aggShape() *shape {
	for len(sh.queuePos) > 0 {
		p := sh.queuePos[0]
		sh.queuePos = sh.queuePos[1:]

		if _, okInt := sh.intern[p]; okInt {
			continue
		}

		deltas := [][2]int{
			{-1, 0},
			{1, 0},
			{0, 1},
			{0, -1},
		}

		validNext := make([]point, 0, len(deltas))

		for i := range deltas {
			dr, dc := p.r+deltas[i][0], p.c+deltas[i][1]
			validNext = append(validNext, point{dr, dc, sh.mtx[dr][dc]})
		}

		nextInt := lo.Filter(validNext, func(np point, _ int) bool { return np.char == p.char })
		nextExt := lo.Filter(validNext, func(np point, _ int) bool { return np.char != p.char })

		if len(nextExt) > 0 {
			sh.border[p]++

			for _, extPos := range nextExt {
				sh.outer[extPos]++
			}
		}

		sh.intern[p]++
		if len(nextInt) > 0 {
			sh.queuePos = append(sh.queuePos, nextInt...)
		}
	}

	return sh
}

func (sh *shape) countSides() int {
	qLines := make([]line, 0, len(sh.border))

	for borderPoint := range sh.border {
		qLines = append(qLines, sh.pointToLines(borderPoint)...)
	}

	for len(qLines) > 0 {
		curLine := qLines[0]
		qLines = qLines[1:]

		if _, visited := sh.visited[curLine.e1]; visited {
			continue
		}

		curLine.aggLine(sh.mtx)

		sh.addVisited(curLine.visited)
		sh.lines = append(sh.lines, curLine)
	}

	return len(sh.lines)
}

func (sh *shape) addVisited(in map[dirPoint]int) {
	for k, v := range in {
		sh.visited[k] += v
	}
}

func (sh *shape) pointToLines(p point) []line {
	lines := []line{}

	if p.char == 0 {
		return lines
	}

	deltas := map[dir][2]int{
		UP:    {-1, 0},
		DOWN:  {1, 0},
		RIGHT: {0, 1},
		LEFT:  {0, -1},
	}

	for dir, delta := range deltas {
		newP := newDPoint(sh.mtx, p, delta, dir)
		if p.char != newP.char {
			dp := dirPoint{p.r, p.c, p.char, dir}
			lines = append(lines, line{dp, dp, dir, make(map[dirPoint]int)})
		}
	}

	return lines
}

type shape struct {
	id       int
	queuePos []point
	mtx      [][]byte
	intern   map[point]int
	border   map[point]int
	outer    map[point]int
	lines    []line
	visited  map[dirPoint]int
}

type dir string

const (
	UP    dir = "^"
	DOWN  dir = "v"
	LEFT  dir = "<"
	RIGHT dir = ">"
)

type line struct {
	e1, e2    dirPoint
	normalDir dir
	visited   map[dirPoint]int
}

func (l *line) aggLine(mtx [][]byte) {
	l.visited[l.e1]++
	l.visited[l.e2]++

	switch l.normalDir {
	case UP, DOWN:
		l.extendHorizontal(mtx)
	case LEFT, RIGHT:
		l.extentVertical(mtx)
	}
}

func (l *line) extendHorizontal(mtx [][]byte) {
	diagDelta := -1
	newE1, newE2 := l.e1, l.e2
	if l.normalDir == DOWN {
		diagDelta = 1
	}

	//Right
	i := 0
	for l.e2.c+i < len(mtx[0]) && mtx[l.e2.r][l.e2.c+i] == l.e2.char && mtx[l.e2.r+diagDelta][l.e2.c+i] != l.e2.char {
		newE2 = dirPoint{l.e2.r, l.e2.c + i, l.e2.char, l.e2.dir}
		l.addVisited(newE2)
		i++
	}

	//Left
	i = 0
	for l.e1.c-i >= 0 && mtx[l.e1.r][l.e1.c-i] == l.e1.char && mtx[l.e1.r+diagDelta][l.e2.c-i] != l.e1.char {
		newE1 = dirPoint{l.e1.r, l.e1.c - i, l.e1.char, l.e1.dir}
		l.addVisited(newE1)
		i++
	}

	l.e1, l.e2 = newE1, newE2
}

func (l *line) extentVertical(mtx [][]byte) {
	diagDelta := -1
	newE1, newE2 := l.e1, l.e2
	if l.normalDir == RIGHT {
		diagDelta = 1
	}

	//Down
	i := 0
	for l.e2.r+i < len(mtx) && mtx[l.e2.r+i][l.e2.c] == l.e2.char && mtx[l.e2.r+i][l.e2.c+diagDelta] != l.e2.char {
		newE2 = dirPoint{l.e2.r + i, l.e2.c, l.e2.char, l.e2.dir}
		l.addVisited(newE2)
		i++
	}

	//Up
	i = 0
	for l.e1.r-i >= 0 && mtx[l.e1.r-i][l.e1.c] == l.e1.char && mtx[l.e2.r-i][l.e1.c+diagDelta] != l.e1.char {
		newE1 = dirPoint{l.e1.r - i, l.e1.c, l.e1.char, l.e1.dir}
		l.addVisited(newE1)
		i++
	}

	l.e1, l.e2 = newE1, newE2
}

func (l *line) addVisited(p dirPoint) {
	l.visited[p]++
}

type point struct {
	r, c int
	char byte
}

type dirPoint struct {
	r, c int
	char byte
	dir  dir
}

func newDPoint(mtx [][]byte, p point, delta [2]int, dir dir) dirPoint {
	r, c := p.r+delta[0], p.c+delta[1]
	return dirPoint{r, c, mtx[r][c], dir}
}
