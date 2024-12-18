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
	slog.Info("Part 1", "time", time.Since(start).String(), "Ans", new(solver).prep(maxGrid, bytesN, inp).part1(maxGrid))

	start2 := time.Now()
	slog.Info("Part 2", "time", time.Since(start2).String(), "Ans", part2(maxGrid, bytesN, inp))
}

type solver struct {
	mtx    [][]uint8
	broken []coord
}

func (s *solver) prep(space, n int, inp string) *solver {
	maxCoord := space + 1
	s.mtx = make([][]uint8, maxCoord)
	for i := range maxCoord {
		s.mtx[i] = make([]uint8, maxCoord)
	}

	for i, pos := range strings.Split(inp, "\n") {
		point := strings.Split(pos, ",")
		x, _ := strconv.ParseInt(point[0], 10, 8)
		y, _ := strconv.ParseInt(point[1], 10, 8)
		s.broken = append(s.broken, newCoord(int(y), int(x)))
		if i < n {
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
		curDPos := q.get()
		if curDPos.score > bestScore {
			continue
		}

		curPos := curDPos.dpos
		visited[curPos.hash()] = curDPos.score

		if curDPos.dpos.eqPos(end) {
			if curDPos.score < bestScore {
				bestScore = curDPos.score
			}

			winners = append(winners, curDPos)
			continue
		}

		opts := options(visited, end, curDPos, s.mtx)
		for o := range opts {
			q.push(opts[o])
		}

		q.sortAsc()
	}

	minpath := math.MaxInt
	for _, w := range winners {
		if len(w.path) < minpath {
			minpath = len(w.path) - 1
		}
	}

	return minpath
}

func part2(maxGrid, bytesN int, inp string) string {
	s0 := new(solver).prep(maxGrid, bytesN, inp)
	dists := []res{}
	testSet := s0.broken[bytesN:]
	outs := make(chan res, len(testSet))

	wg := sync.WaitGroup{}
	wg.Add(len(testSet))

	for i, bcor := range testSet {
		go func(i int, bcoor coord) {
			defer wg.Done()

			s := new(solver).prep(maxGrid, bytesN, inp)
			s.mtx[bcoor.r][bcoor.c] = 1

			got := s.part1(maxGrid)
			blockRes := res{
				i:    i,
				dist: got,
				pos:  bcoor,
				out:  bcoor.str(),
			}

			dists = append(dists, blockRes)

			if got == math.MaxInt {
				outs <- blockRes
			}
		}(i, bcor)
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

	fmt.Println(dists)
	fmt.Println(minCoord)

	return minCoord.str()
}

type res struct {
	i    int
	dist int
	pos  coord
	out  string
}

func startOptions(start, end coord, mtx [][]uint8) []step {
	stE := step{
		parent: newDCoord(start, EAST),
		dpos:   newDCoord(start, EAST),
		score:  0,
		path:   []coord{},
	}
	stS := step{
		parent: newDCoord(start, SOUTH),
		dpos:   newDCoord(start, SOUTH),
		score:  0,
		path:   []coord{},
	}

	return append(options(map[string]int{}, end, stS, mtx), options(map[string]int{}, end, stE, mtx)...)
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

func options(visited map[string]int, target coord, curStep step, mtx [][]uint8) []step {
	out := []step{}
	for dir := range deltas {
		row, col := curStep.dpos.c.r+deltas[dir][0], curStep.dpos.c.c+deltas[dir][1]
		if dir == opposites[curStep.dpos.dir] ||
			row < 0 || col < 0 ||
			row >= len(mtx) || col >= len(mtx[0]) ||
			mtx[row][col] == 1 {
			continue
		}

		newPath := make([]coord, len(curStep.path))
		copy(newPath, curStep.path)

		stepCoord := newCoord(row, col)
		stepDirCoord := newDCoord(stepCoord, dir)
		newScore := len(newPath) + manhattanDistance(stepCoord, target)

		st := step{
			parent: curStep.dpos,
			dpos:   stepDirCoord,
			score:  newScore,
			path:   append(newPath, stepCoord),
		}

		wasScore, wasVisited := visited[stepDirCoord.hash()]
		if wasVisited && wasScore <= newScore {
			continue
		}

		out = append(out, st)
	}

	curStep.path = nil

	return out
}

func manhattanDistance(a, b coord) int {
	ay, ax := float64(a.r), float64(a.c)
	by, bx := float64(b.r), float64(b.c)
	return int(math.Abs(ax-bx) + math.Abs(ay-by))
}
