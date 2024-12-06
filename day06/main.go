package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log/slog"
	"strings"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Pos:", new(solver).prep(inp).part1())
	slog.Info("Part 2:", "Pos:", new(solver).prep(inp).part2(inp))
}

type pos struct {
	r, c int
}

type dpos struct {
	dir      dir
	row, col int
}

func newPos(r, c int) pos {
	return pos{r: r, c: c}
}

func newDirPos(r, c int, d dir) dpos {
	return dpos{row: r, col: c, dir: d}
}

type dir string
type cmd int

const (
	obst = byte('#')

	UP    dir = "^"
	RIGHT dir = ">"
	DOWN  dir = "v"
	LEFT  dir = "<"

	STOP    cmd = -1
	LOOPED  cmd = -2
	NEW_DIR cmd = 0
	KEEP    cmd = 1
)

var nextDir = map[dir]dir{
	UP:    RIGHT,
	RIGHT: DOWN,
	DOWN:  LEFT,
	LEFT:  UP,
}

type solver struct {
	mtx   [][]byte
	guard dpos
	init  struct {
		guard dpos
	}
	visited    map[pos]int
	dirVisited map[dpos]int
}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n")
	s.mtx = make([][]byte, 0, len(rows))
	s.visited = make(map[pos]int)
	s.dirVisited = make(map[dpos]int)

	for i := range rows {
		newRow := []byte(rows[i])
		if startCol := bytes.IndexAny(newRow, "^><v"); startCol != -1 {
			s.visit(i, startCol, dir(newRow[startCol]))
			s.init.guard = newDirPos(i, startCol, dir(newRow[startCol]))
			newRow[startCol] = '.'
		}
		s.mtx = append(s.mtx, newRow)
	}

	return s
}

func (s *solver) part1() int {
	res := KEEP
	for res != STOP {
		switch s.guard.dir {
		case UP, DOWN:
			res = s.vertical()
		case LEFT, RIGHT:
			res = s.horizontal()
		}
	}
	return len(s.visited)
}

func (s *solver) part2(inp string) int {
	s0 := new(solver).prep(inp)
	s0.part1()

	loopedPos := make(map[pos]int)
	options := s.getOptions(s0.visited)
	// options := s.getTruncOptions(s0.dirVisited)

	for _, o := range options {
		s.mtx[o.r][o.c] = obst
		res := KEEP

		for res != STOP && res != LOOPED {
			switch s.guard.dir {
			case UP, DOWN:
				res = s.vertical()
			case LEFT, RIGHT:
				res = s.horizontal()
			}
		}

		if res == LOOPED {
			loopedPos[newPos(o.r, o.c)] += 1
		}

		s.reset(o)
	}

	return len(loopedPos)
}

func (s *solver) vertical() cmd {
	c := s.guard.col
	start, step := 0, 0

	if s.guard.dir == UP {
		start = s.guard.row - 1
		step = -1
	}
	if s.guard.dir == DOWN {
		start = s.guard.row
		step = 1
	}

	for r := start; s.inBound(r, len(s.mtx)); r += step {
		if res := s.checkUpdate(r, c); res != KEEP {
			return res
		}
	}

	return STOP
}

func (s *solver) horizontal() cmd {
	r := s.guard.row
	start, step := 0, 0

	if s.guard.dir == LEFT {
		start = s.guard.col - 1
		step = -1
	}
	if s.guard.dir == RIGHT {
		start = s.guard.col
		step = 1
	}

	for c := start; s.inBound(c, len(s.mtx[r])); c += step {
		if res := s.checkUpdate(r, c); res != KEEP {
			return res
		}
	}

	return STOP
}

func (s *solver) inBound(val, max int) bool {
	return val > -1 && val < max
}

func (s *solver) checkUpdate(r, c int) cmd {
	if _, ok := s.dirVisited[newDirPos(r, c, s.guard.dir)]; ok {
		return LOOPED
	}
	if s.mtx[r][c] == obst {
		s.guard.dir = nextDir[s.guard.dir]
		return NEW_DIR
	}
	s.visit(r, c, s.guard.dir)

	return KEEP
}

func (s *solver) visit(r, c int, d dir) {
	s.visited[newPos(r, c)] += 1
	s.dirVisited[newDirPos(r, c, d)] += 1
	s.guard = newDirPos(r, c, d)
}

func (s *solver) reset(o pos) {
	s.guard = s.init.guard

	s.dirVisited = make(map[dpos]int)
	s.visited = make(map[pos]int)

	s.visit(s.guard.row, s.guard.col, s.guard.dir)

	s.mtx[o.r][o.c] = '.'
}

// returned 5.145 options
func (s *solver) getOptions(visited map[pos]int) []pos {
	out := make([]pos, 0, len(visited))

	for p := range visited {
		out = append(out, p)
	}

	return out
}

// getTruncOptions is not working properly
// idea - limit options to only those, moving from which prev cell would result in previously visited path
// returned 1.371 option
func (s *solver) getTruncOptions(s0dirVisited map[dpos]int) []pos {
	out := make([]pos, 0, len(s0dirVisited))
	start := newDirPos(s.init.guard.row, s.init.guard.col, s.init.guard.dir)

	for point := range s0dirVisited {
		if point == start {
			continue
		}

		if s.isPromissing(point, s0dirVisited) {
			out = append(out, newPos(point.row, point.col))
		}
	}

	return out
}

func (s *solver) isPromissing(option dpos, s0dirVisited map[dpos]int) bool {
	stepToPrev := [2]int{0, 0}
	dSearchOptionMove := [2]int{0, 0}

	switch option.dir {
	case UP:
		stepToPrev = [2]int{1, 0}
		dSearchOptionMove = [2]int{0, 1}
	case DOWN:
		stepToPrev = [2]int{-1, 0}
		dSearchOptionMove = [2]int{0, -1}
	case LEFT:
		stepToPrev = [2]int{0, 1}
		dSearchOptionMove = [2]int{-1, 0}
	case RIGHT:
		stepToPrev = [2]int{0, -1}
		dSearchOptionMove = [2]int{1, 0}
	}

	wantDir := nextDir[option.dir]
	prev := newDirPos(option.row+stepToPrev[0], option.col+stepToPrev[1], option.dir)
	r, c := prev.row, prev.col

	for r > -1 && r < len(s.mtx) && c > -1 && c < len(s.mtx[r]) {
		checkPoint := newDirPos(r, c, wantDir)
		if _, ok := s0dirVisited[checkPoint]; ok {
			return true
		}

		r, c = r+dSearchOptionMove[0], c+dSearchOptionMove[1]
	}

	return false
}

func (s *solver) print() {
	for r := range s.mtx {
		prow := s.mtx[r]
		for k := range s.visited {
			if k.r == r {
				prow[k.c] = 'X'
			}
		}
		fmt.Println(string(prow))
	}
	fmt.Println()
}

func (s *solver) printBlocked(blocked map[pos]int) {
	for r := range s.mtx {
		prow := s.mtx[r]
		for k := range blocked {
			if k.r == r {
				prow[k.c] = 'O'
			}
		}
		fmt.Println(string(prow))
	}
	fmt.Println()
}
