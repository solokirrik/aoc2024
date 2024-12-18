package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
	"sync"
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

func (s *solver) part1(maxCoord int) int {
	q := newQQ()
	start := coord{maxCoord, maxCoord}
	end := coord{0, 0}
	bestScore := math.MaxInt
	winners := make([]step, 0, 20)
	visited := make(map[string]int, len(s.mtx[0])*len(s.mtx))

	opts := startOptions(start, end, s.mtx)
	for _, o := range opts {
		q.push(o)
	}

	for q.len() > 0 {
		curStep := q.get()
		if curStep.score > bestScore {
			continue
		}

		curDPos := curStep.dpos
		wasScore, wasVisited := visited[curDPos.hash()]
		if wasVisited && wasScore <= curStep.score {
			continue
		}

		visited[curDPos.hash()] = curStep.score

		if curStep.dpos.eqPos(end) {
			if curStep.score < bestScore {
				bestScore = curStep.score
			}

			winners = append(winners, curStep)
			continue
		}

		opts := options(visited, curStep, s.mtx)
		for o := range opts {
			q.push(opts[o])
		}

		q.sortAsc()
	}

	minpath := math.MaxInt
	for _, w := range winners {
		if w.pathLen < minpath {
			minpath = w.pathLen
		}
	}

	return minpath
}

func part2(maxGrid, brokenBytes int, inp string) string {
	s0 := new(solver).prep(maxGrid, brokenBytes, inp)
	outs := make(chan res, len(s0.broken)-brokenBytes)

	wg := sync.WaitGroup{}
	wg.Add(len(s0.broken) - brokenBytes)

	for i := brokenBytes; i < len(s0.broken); i++ {
		go func(i int, bcoor coord) {
			defer wg.Done()

			s := new(solver).prep(maxGrid, i, inp)
			s.mtx[bcoor.r][bcoor.c] = 1

			got := s.part1(maxGrid)
			blockRes := res{
				i:    i,
				dist: got,
				pos:  bcoor,
			}

			if got == math.MaxInt {
				outs <- blockRes
			}
		}(i, s0.broken[i])
	}

	wg.Wait()

	close(outs)

	minCoord := coord{}
	minI := math.MaxInt

	for val := range outs {
		fmt.Println(val)
		if val.i < minI {
			minI = val.i
			minCoord = coord{r: val.pos.r, c: val.pos.c}
		}
	}

	return minCoord.str()
}

type res struct {
	i    int
	dist int
	pos  coord
}

func startOptions(start, end coord, mtx [][]uint8) []step {
	stE := step{
		parent: newDCoord(start, EAST),
		dpos:   newDCoord(start, EAST),
	}
	stS := step{
		parent: newDCoord(start, SOUTH),
		dpos:   newDCoord(start, SOUTH),
	}

	return append(
		options(map[string]int{}, stS, mtx),
		options(map[string]int{}, stE, mtx)...,
	)
}

const (
	NORTH = "^"
	EAST  = ">"
	SOUTH = "v"
	WEST  = "<"
)

var (
	deltas = map[string][2]int{
		NORTH: {-1, 0},
		EAST:  {0, 1},
		SOUTH: {1, 0},
		WEST:  {0, -1},
	}

	opposites = map[string]string{
		NORTH: SOUTH,
		EAST:  WEST,
		SOUTH: NORTH,
		WEST:  EAST,
	}
)

func options(visited map[string]int, curStep step, mtx [][]uint8) []step {
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

		wasScore, wasVisited := visited[stepDirCoord.hash()]
		if wasVisited && wasScore <= newScore {
			continue
		}

		out = append(out, st)
	}

	return out
}
