package main

import (
	_ "embed"
	"image"
	"log/slog"
	"math"
	"strings"
	"time"
)

//go:embed inp
var inp string

func main() {
	start1 := time.Now()
	got1 := new(solver).prep(inp).part1()
	slog.Info("Part 1:", "time", time.Since(start1).String(), "Ans", got1)

	start2 := time.Now()
	got2 := new(solver).prep(inp).part2()
	slog.Info("Part 2:", "time", time.Since(start2).String(), "Ans", got2)
}

type solver struct {
	mtx     [][]string
	start   dcoord
	end     coord
	visited map[string]int
	doVisio bool
}

const (
	NORTH = "^"
	EAST  = ">"
	SOUTH = "v"
	WEST  = "<"
)

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n")
	s.mtx = make([][]string, 0, len(rows))

	for r, raw := range rows {
		row := strings.Split(raw, "")

		if start := strings.Index(raw, "S"); start != -1 {
			s.start = newDCoord(newCoord(r, start), EAST)
			row[start] = "."
		}

		if end := strings.Index(raw, "E"); end != -1 {
			s.end = newCoord(r, end)
			row[end] = "."
		}

		s.mtx = append(s.mtx, row)
	}

	return s
}

func (s *solver) part1() int {
	q := newQQ()
	bestScore := math.MaxInt
	winners := make([]step, 0, 20)
	s.visited = make(map[string]int, len(s.mtx[0])*len(s.mtx))

	q.push(step{
		score:  0,
		pos:    s.start,
		parent: s.start,
		path:   []coord{s.start.getPos()},
	})

	var images []*image.Paletted
	var delays []int

	i := -1
	for q.len() > 0 {
		i++
		if s.doVisio && i%1000 == 0 {
			frame := s.visitedFrame()
			images = append(images, frame)
			delays = append(delays, 5)
		}

		curPos := q.get()
		if curPos.score > bestScore {
			continue
		}

		s.visited[curPos.pos.id] = curPos.score

		if curPos.pos.eqPos(s.end) {
			if curPos.score < bestScore {
				bestScore = curPos.score
			}

			winners = append(winners, curPos)
			continue
		}

		opts := options(curPos, s.mtx)
		for o := range opts {
			if visScore, ok := s.visited[opts[o].pos.id]; ok && visScore < opts[o].score {
				continue
			}
			q.push(opts[o])
		}

		q.sort()
	}

	if s.doVisio {
		s.writeGif(images, delays)
		s.savePaths(bestScore, winners)
	}

	return bestScore
}

func (s *solver) part2() int {
	q := newQQ()
	bestScore := math.MaxInt
	winners := make([]step, 0, 20)
	s.visited = make(map[string]int, len(s.mtx[0])*len(s.mtx))

	q.push(step{
		score:  0,
		pos:    s.start,
		parent: s.start,
		path:   []coord{s.start.getPos()},
	})

	for q.len() > 0 {
		curPos := q.get()
		if curPos.score > bestScore {
			continue
		}

		s.visited[curPos.pos.id] = curPos.score

		if curPos.pos.eqPos(s.end) {
			if curPos.score < bestScore {
				bestScore = curPos.score
			}

			winners = append(winners, curPos)
			continue
		}

		opts := options(curPos, s.mtx)
		for o := range opts {
			if visScore, ok := s.visited[opts[o].pos.id]; !ok || opts[o].score < visScore {
				q.push(opts[o])
			}
		}

		q.sort()
	}

	return overlappingSeats(bestScore, winners)
}

func overlappingSeats(bestScore int, winners []step) int {
	paths := make([]coord, 0, len(winners)*len(winners[0].path))
	for w := range winners {
		if winners[w].score != bestScore {
			continue
		}

		paths = append(paths, winners[w].path...)
	}

	seats := make(map[string]coord, len(paths))
	for i := range paths {
		seats[paths[i].hash()] = paths[i]
	}

	return len(seats)
}

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

func options(curStep step, mtx [][]string) []step {
	out := []step{}
	for dir := range deltas {
		row, col := curStep.pos.c.r+deltas[dir][0], curStep.pos.c.c+deltas[dir][1]
		if dir == opposites[curStep.pos.dir] ||
			row < 0 || col < 0 ||
			row >= len(mtx) || col >= len(mtx[0]) ||
			mtx[row][col] == "#" {
			continue
		}

		newPath := make([]coord, len(curStep.path))
		copy(newPath, curStep.path)

		st := step{
			parent: curStep.pos,
			pos:    newDCoord(newCoord(row, col), dir),
			score:  curStep.score + score(curStep.pos, dir),
			path:   append(newPath, newCoord(row, col)),
		}
		out = append(out, st)
	}

	curStep.path = nil

	return out
}

func score(parent dcoord, dir string) int {
	if parent.dir == dir {
		return 1
	}

	return 1001
}
