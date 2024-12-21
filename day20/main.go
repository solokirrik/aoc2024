package main

import (
	_ "embed"
	"log/slog"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/solokirrik/aoc2024/pkg"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Starting day20")
	start := time.Now()

	got1 := new(solver).prep(inp).part1(100)
	slog.Info("Part 1", "time", time.Since(start).String(), "Ans", got1)

	start2 := time.Now()
	got2 := new(solver).prep(inp).part2(100)
	slog.Info("Part 2", "time", time.Since(start2).String(), "Ans", got2)
}

type solver struct {
	mtx        [][]string
	start, end pkg.Coord
}

func (s *solver) prep(inp string) *solver {
	rows := strings.Split(inp, "\n")
	s.mtx = make([][]string, 0, len(rows))

	for i := range rows {
		s.mtx = append(s.mtx, strings.Split(rows[i], ""))
		if idx := strings.Index(rows[i], "S"); idx != -1 {
			s.start = pkg.NewCoord(i, idx)
		}
		if idx := strings.Index(rows[i], "E"); idx != -1 {
			s.end = pkg.NewCoord(i, idx)
		}
	}

	return s
}

const minDiff = 100

func (s *solver) part1(minDiff int) int {
	basePath := solve(s.mtx, s.start, s.end, true)
	cheats := partCheatsP1(basePath, minDiff)

	return len(cheats)
}

func (s *solver) part2(minDiff int) int {
	basePath := solve(s.mtx, s.start, s.end, true)
	cheats := partCheatsP2(s.mtx, basePath, minDiff)

	return len(cheats)
}

func partCheatsP1(path []pkg.Coord, minDiff int) map[[2]pkg.Coord]bool {
	out := make(map[[2]pkg.Coord]bool)

	for pointIndex, pathPoint := range path {
		for dir, delta := range pkg.Deltas {
			dR, dC := delta[0]*2, delta[1]*2
			checkPoint := pkg.NewCoord(pathPoint.R+dR, pathPoint.C+dC)

			idx := slices.Index(path, pkg.NewCoord(pathPoint.R+dR, pathPoint.C+dC))
			if idx-pointIndex < minDiff+1 {
				continue
			}

			pair := [2]pkg.Coord{
				pkg.NewCoord(pathPoint.R+delta[0], pathPoint.C+delta[1]),
				checkPoint,
			}

			if dir == pkg.SOUTH || dir == pkg.NORTH {
				pair[0].R, pair[1].R = min(pair[0].R, pair[1].R), max(pair[0].R, pair[1].R)
			}
			if dir == pkg.EAST || dir == pkg.WEST {
				pair[0].C, pair[1].C = min(pair[0].C, pair[1].C), max(pair[0].C, pair[1].C)
			}
			out[pair] = true
		}
	}

	return out
}

func partCheatsP2(mtx [][]string, path []pkg.Coord, minDiff int) map[[2]pkg.Coord]bool {
	out := make(map[[2]pkg.Coord]bool)

	for pointSIndex, pathSPoint := range path {
		for _, delta := range pkg.Deltas {
			dR, dC := delta[0], delta[1]
			checkPoint := pkg.NewCoord(pathSPoint.R+dR, pathSPoint.C+dC)
			if mtx[checkPoint.R][checkPoint.C] != "#" {
				continue
			}

			for pointEIndex, pathEPoint := range path {
				if pointSIndex == pointEIndex {
					continue
				}

				if manhattanDistance(checkPoint, pathEPoint) <= 20-1 &&
					pointEIndex-pointSIndex >= minDiff {
					pair := [2]pkg.Coord{
						pathSPoint,
						pathEPoint,
					}

					out[pair] = true
				}
			}
			break
		}
	}

	return out
}

const NONE = -1

func solve(mtx [][]string, start, end pkg.Coord, savePath bool) []pkg.Coord {
	qq := pkg.NewStepQueue()
	visited := make(map[pkg.CoordHash]int, len(mtx[0])*len(mtx))

	opts := startOptions(start, mtx, savePath)
	for _, o := range opts {
		qq.Push(o)
	}

	for qq.Len() > 0 {
		curStep := qq.Get()
		curPos := curStep.DPos.GetPos()
		if _, wasVisited := visited[curPos.Hash()]; wasVisited {
			continue
		}

		visited[curPos.Hash()] = curStep.Score

		if curPos.Eq(end) {
			if !savePath {
				return make([]pkg.Coord, curStep.PathLen)
			}
			return curStep.Path
		}

		opts := options(visited, curStep, mtx, savePath)
		for o := range opts {
			qq.Push(opts[o])
		}
	}

	return []pkg.Coord{}
}

func startOptions(start pkg.Coord, mtx [][]string, savePath bool) []pkg.Step {
	out := []pkg.Step{}
	for dir := range pkg.Deltas {
		st := pkg.Step{
			Parent: pkg.NewDCoord(start, dir),
			DPos:   pkg.NewDCoord(start, dir),
		}
		out = append(out, options(map[pkg.CoordHash]int{}, st, mtx, savePath)...)
	}

	return out
}

func options(
	visited map[pkg.CoordHash]int,
	curStep pkg.Step,
	mtx [][]string,
	savePath bool,
) []pkg.Step {
	out := []pkg.Step{}
	for dir := range pkg.Deltas {
		row, col := curStep.DPos.C.R+pkg.Deltas[dir][0], curStep.DPos.C.C+pkg.Deltas[dir][1]
		if dir == pkg.Opposites[curStep.DPos.Dir] ||
			row < 0 || col < 0 ||
			row >= len(mtx) || col >= len(mtx[0]) ||
			mtx[row][col] == "#" {
			continue
		}

		stepCoord := pkg.NewCoord(row, col)
		stepDirCoord := pkg.NewDCoord(stepCoord, dir)
		newScore := curStep.PathLen

		var newPath []pkg.Coord
		if savePath {
			newPath = append(curStep.Path, stepCoord)
		}

		st := pkg.Step{
			Parent:  curStep.DPos,
			DPos:    stepDirCoord,
			Score:   newScore,
			PathLen: curStep.PathLen + 1,
			Path:    newPath,
		}

		if _, wasVisited := visited[stepCoord.Hash()]; wasVisited {
			continue
		}

		out = append(out, st)
	}

	return out
}

func manhattanDistance(a, b pkg.Coord) int {
	ay, ax := float64(a.R), float64(a.C)
	by, bx := float64(b.R), float64(b.C)
	return int(math.Abs(ax-bx) + math.Abs(ay-by))
}
