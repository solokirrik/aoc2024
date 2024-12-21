package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/solokirrik/aoc2024/pkg"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Starting day")

	start := time.Now()
	got1 := new(solver).prep(inp).part1()
	slog.Info("Part 1", "time", time.Since(start).String(), "Ans", got1)

	start2 := time.Now()
	got2 := new(solver).prep(inp).part2()
	slog.Info("Part 2", "time", time.Since(start2).String(), "Ans", got2)
}

const NONE = "#"

var (
	numpad = [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{NONE, "0", "A"},
	}

	arrowpad = [][]string{
		{NONE, "^", "A"},
		{"<", "v", ">"},
	}

	numALoc   = pkg.Coord{R: 3, C: 2}
	arrowALoc = pkg.Coord{R: 0, C: 2}

	numCoord = map[string]pkg.Coord{
		"1": {R: 2, C: 0}, "2": {R: 2, C: 1}, "3": {R: 2, C: 2},
		"4": {R: 1, C: 0}, "5": {R: 1, C: 1}, "6": {R: 1, C: 2},
		"7": {R: 0, C: 0}, "8": {R: 0, C: 1}, "9": {R: 0, C: 2},
		NONE: {R: 3, C: 0}, "0": {R: 3, C: 1}, "A": {R: 3, C: 2},
	}
	arrowCoord = map[string]pkg.Coord{
		NONE: {R: 0, C: 0}, "^": {R: 0, C: 1}, "A": {R: 0, C: 2},
		"<": {R: 1, C: 0}, "v": {R: 1, C: 1}, ">": {R: 1, C: 2},
	}

	dCoordArrowMap = map[[2]int]string{
		{0, 0}:  "A",
		{-1, 0}: "^",
		{0, 1}:  ">",
		{1, 0}:  "v",
		{0, -1}: "<",
	}
)

type solver struct {
	codes []string
}

func (s *solver) prep(inp string) *solver {
	s.codes = strings.Split(inp, "\n")

	return s
}

func (s *solver) part1() int {
	res := 0

	for c := range s.codes {
		code := s.codes[c]
		path1 := numToArrows(code)
		path2 := arrToArr(path1)
		val, err := strconv.ParseInt(code[:len(code)-1], 10, 64)
		pkg.PanicOnErr(err)

		resLoc := int(val) * len(path2)
		res += resLoc
	}

	return res
}

func (s *solver) part2() int {
	return 0
}

func numToArrows(code string) []string {
	out := make([]string, 0, len(code))

	for ch := 0; ch < len(code); ch++ {
		var start, end pkg.Coord

		if ch == 0 {
			start = numALoc
			end = numCoord[string(code[ch])]
		} else {
			start = numCoord[string(code[ch-1])]
			end = numCoord[string(code[ch])]
		}

		path := solve(numpad, start, end, true)
		path[len(path)-1] = numCoord[numpad[end.R][end.C]]
		arrows := pathToArrows(path)
		out = append(out, arrows...)
		out = append(out, "A")
	}

	fmt.Println(code, strings.Join(out, ""))

	return out
}

func arrToArr(path []string) []string {
	out := make([]string, 0, len(path))

	for ch := 0; ch < len(path); ch++ {
		var start, end pkg.Coord

		if ch == 0 {
			start = arrowALoc
			end = arrowCoord[path[ch]]
		} else {
			start = arrowCoord[path[ch-1]]
			end = arrowCoord[path[ch]]
		}

		path := solve(arrowpad, start, end, true)
		path[len(path)-1] = numCoord[numpad[end.R][end.C]]
		arrows := pathToArrows(path)
		out = append(out, arrows...)
		out = append(out, "A")
		fmt.Println(strings.Join(out, ""))
	}

	fmt.Println(strings.Join(path, ""), strings.Join(out, ""))

	return out
}

func pathToArrows(path []pkg.Coord) []string {
	out := make([]string, 0, len(path)-1)
	for i := 1; i < len(path); i++ {
		out = append(
			out,
			dCoordArrowMap[[2]int{
				path[i].R - path[i-1].R,
				path[i].C - path[i-1].C,
			}])
	}

	return out
}

func solve(mtx [][]string, start, end pkg.Coord, savePath bool) []pkg.Coord {
	if start.Eq(end) {
		return []pkg.Coord{start, end}
	}

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
			Path:   []pkg.Coord{start},
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
