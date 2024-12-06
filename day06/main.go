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
	slog.Info("Part 2:", "Pos:", new(solver).prep(inp).part2())
}

type pos struct {
	r, c int
}

type dir string

const (
	UP    = dir("^")
	RIGHT = dir(">")
	DOWN  = dir("v")
	LEFT  = dir("<")
	STOP  = -1
)

type solver struct {
	guard   pos
	dir     dir
	visited map[pos]int
	mtx     [][]byte
}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n")
	s.mtx = make([][]byte, 0, len(rows))
	s.visited = make(map[pos]int)

	for i := range rows {
		newRow := []byte(rows[i])
		if startCol := bytes.IndexAny(newRow, "^><v"); startCol != -1 {
			s.visit(i, startCol)
			s.dir = dir(newRow[s.guard.c])
			newRow[startCol] = '.'
		}
		s.mtx = append(s.mtx, newRow)
	}

	return s
}

var obst = byte('#')

func (s *solver) part1() int {
	res := 0
	for res != STOP {
		switch s.dir {
		case UP:
			res = s.decideOnUp()
		case RIGHT:
			res = s.decideOnRight()
		case DOWN:
			res = s.decideOnDown()
		case LEFT:
			res = s.decideOnLeft()
		}
	}
	return len(s.visited)
}

func (s *solver) decideOnUp() int {
	c := s.guard.c
	for r := s.guard.r - 1; r > -1; r-- {
		if s.mtx[r][c] == obst {
			s.dir = RIGHT
			return 0
		}
		s.visit(r, c)
	}
	return STOP
}

func (s *solver) decideOnRight() int {
	r := s.guard.r
	for c := s.guard.c + 1; c < len(s.mtx[r]); c++ {
		if s.mtx[r][c] == obst {
			s.dir = DOWN
			return 0
		}
		s.visit(r, c)
	}
	return STOP
}

func (s *solver) decideOnDown() int {
	c := s.guard.c
	for r := s.guard.r; r < len(s.mtx); r++ {
		if s.mtx[r][c] == obst {
			s.dir = LEFT
			return 0
		}
		s.visit(r, c)
	}
	return STOP
}

func (s *solver) decideOnLeft() int {
	r := s.guard.r
	for c := s.guard.c - 1; c > -1; c-- {
		if s.mtx[r][c] == obst {
			s.dir = UP
			return 0
		}
		s.visit(r, c)
		s.guard = pos{r, c}
	}
	return STOP
}

func (s *solver) visit(r, c int) {
	s.visited[pos{r, c}] += 1
	s.guard = pos{r, c}
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

func (s *solver) part2() int {
	sum := 0

	return sum
}
