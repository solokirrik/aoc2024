package main

import (
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

//go:embed inp
var inp string

//go:embed ex
var ex string

const (
	maxGrid   = 70
	bytesN    = 1024
	maxGridEx = 6
	bytesEx   = 12
)

func main() {
	slog.Info("Starting day18")
	start := time.Now()
	got1 := new(solver).prep(maxGrid, bytesN, inp).part1(maxGrid)
	slog.Info("Part 1", "time", time.Since(start).String(), "Ans", got1)

	start2 := time.Now()
	got2 := part2(maxGrid, bytesN, inp)
	slog.Info("Part 2", "time", time.Since(start2).String(), "Ans", got2)
}

type solver struct {
	mtx    [][]uint8
	broken []coord
}

func (s *solver) prep(maxGrid, brokenBytes int, inp string) *solver {
	maxCoord := maxGrid + 1
	s.mtx = make([][]uint8, maxCoord)

	for i := range maxCoord {
		s.mtx[i] = make([]uint8, maxCoord)
	}

	for i, pos := range strings.Split(inp, "\n") {
		point := strings.Split(pos, ",")
		x, _ := strconv.ParseInt(point[0], 10, 8)
		y, _ := strconv.ParseInt(point[1], 10, 8)
		s.broken = append(s.broken, newCoord(int(y), int(x)))

		if i <= brokenBytes {
			s.mtx[y][x] = 1
		}
	}

	return s
}

const NONE = -1

func (s *solver) part1(maxCoord int) int {
	q := newQQ()
	start := coord{0, 0}
	end := coord{maxCoord, maxCoord}
	visited := make(map[coordHash]int, len(s.mtx[0])*len(s.mtx))

	opts := startOptions(start, s.mtx)
	for _, o := range opts {
		q.push(o)
	}

	for q.len() > 0 {
		curStep := q.get()
		curPos := curStep.dpos.getPos()
		if _, wasVisited := visited[curPos.hash()]; wasVisited {
			continue
		}

		visited[curPos.hash()] = curStep.score

		if curPos.eq(end) {
			return curStep.pathLen
		}

		opts := options(visited, curStep, s.mtx)
		for o := range opts {
			q.push(opts[o])
		}
	}

	return NONE
}

func part2(maxGrid, brokenBytes int, inp string) string {
	sl := new(solver).prep(maxGrid, brokenBytes, inp)

	for i := brokenBytes; i < len(sl.broken); i++ {
		bcoor := sl.broken[i]
		sl.mtx[bcoor.r][bcoor.c] = 1

		if got := sl.part1(maxGrid); got == NONE {
			return bcoor.str()
		}
	}

	return "NONE"
}

func startOptions(start coord, mtx [][]uint8) []step {
	stE := step{
		parent: newDCoord(start, EAST),
		dpos:   newDCoord(start, EAST),
	}
	stS := step{
		parent: newDCoord(start, SOUTH),
		dpos:   newDCoord(start, SOUTH),
	}

	return append(
		options(map[coordHash]int{}, stS, mtx),
		options(map[coordHash]int{}, stE, mtx)...,
	)
}

const (
	NORTH = 0
	EAST  = 1
	SOUTH = 2
	WEST  = 3
)

var (
	deltas = map[int][2]int{
		NORTH: {-1, 0},
		EAST:  {0, 1},
		SOUTH: {1, 0},
		WEST:  {0, -1},
	}

	opposites = map[int]int{
		NORTH: SOUTH,
		EAST:  WEST,
		SOUTH: NORTH,
		WEST:  EAST,
	}
)

func options(visited map[coordHash]int, curStep step, mtx [][]uint8) []step {
	out := []step{}
	for dir := range deltas {
		row, col := curStep.dpos.c.r+deltas[dir][0], curStep.dpos.c.c+deltas[dir][1]
		if dir == opposites[curStep.dpos.dir] ||
			row < 0 || col < 0 ||
			row >= len(mtx) || col >= len(mtx[0]) ||
			mtx[row][col] == 1 {
			continue
		}

		stepCoord := newCoord(row, col)
		stepDirCoord := newDCoord(stepCoord, dir)
		newScore := curStep.pathLen

		st := step{
			parent:  curStep.dpos,
			dpos:    stepDirCoord,
			score:   newScore,
			pathLen: curStep.pathLen + 1,
		}

		if _, wasVisited := visited[stepCoord.hash()]; wasVisited {
			continue
		}

		out = append(out, st)
	}

	return out
}
