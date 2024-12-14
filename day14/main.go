package main

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/solokirrik/aoc2024/utils"
	"golang.org/x/image/bmp"
)

//go:embed inp
var inp string

const (
	fieldW = 101
	fieldH = 103
)

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp, fieldW, fieldH).part1(100))
	new(solver).prep(inp, fieldW, fieldH).part2(10000, false)
}

type solver struct {
	robots []robot
	fs     coor
}

func (s *solver) prep(inp string, fieldW, fieldH int64) *solver {
	for _, line := range strings.Split(inp, "\n") {
		s.robots = append(s.robots, newRobot(coor{x: fieldW, y: fieldH}, line))
	}

	s.fs = coor{fieldH, fieldH}

	return s
}

func (s *solver) part1(wait int64) int {
	qcount := make(map[int]int)
	for i := range s.robots {
		s.robots[i].move(wait)
		q := s.robots[i].q()
		qcount[q]++
	}

	mul := 1
	for k, v := range qcount {
		if k == 0 {
			continue
		}
		mul *= v
	}

	return mul
}

func (s *solver) part2n(n int64) {
	for i := range s.robots {
		s.robots[i].move(n)
	}
	drawMap(int(n), s.robots)
}

func (s *solver) part2(testSteps int, writeFile bool) {
	var f *os.File
	var err error

	if writeFile {
		f, err = os.Create("./metric_data.csv")
		utils.PanicOnErr(err)

		defer f.Close()

		f.WriteString(fmt.Sprintf("sec,transitions\n"))
	}

	sum := 0
	minV, minIdx := fieldW*fieldH, 0
	maxV, maxIdx := 0, 0
	for i := 1; i < testSteps; i++ {
		for i := range s.robots {
			s.robots[i].move(1)
		}

		tr := countTransitions(s.robots)
		sum += tr
		if tr > maxV {
			maxIdx = i
			maxV = tr
		}
		if tr < minV {
			minIdx = i
			minV = tr
		}

		if writeFile {
			f.WriteString(fmt.Sprintf("%d,%d\n", i, tr))
		}
	}

	slog.Info("Metrics",
		"min_idx", minIdx, "min", minV,
		"max_idx", maxIdx, "max", maxV,
		"avg", sum/testSteps,
	)
}

func countTransitions(rs []robot) int {
	mtx := make([][]byte, fieldH)

	for i := 0; i < fieldH; i++ {
		mtx[i] = make([]uint8, fieldW)
	}

	for _, r := range rs {
		mtx[r.pos.y][r.pos.x] = 1
	}

	transitions := 0
	last := mtx[0][0]

	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			if mtx[i][j] != last {
				last = mtx[i][j]
				transitions++
			}
		}
	}

	return transitions
}

func drawMap(n int, rs []robot) {
	width, height := rs[0].fs.x, rs[0].fs.y
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	for i := 0; i < len(rs); i++ {
		img.Set(int(rs[i].pos.x), int(rs[i].pos.y), color.White)
	}

	fName := "./pics/" + strconv.FormatInt(int64(n), 10) + ".bmp"
	file, err := os.Create(fName)
	utils.PanicOnErr(err)

	defer file.Close()

	utils.PanicOnErr(bmp.Encode(file, img))
}

type robot struct {
	initPos coor
	pos     coor
	vel     coor
	fs      coor
}

type coor struct {
	x, y int64
}

func (r *robot) q() int {
	if r.pos.x < r.fs.x/2 {
		if r.pos.y < r.fs.y/2 {
			return 1
		} else if r.pos.y > r.fs.y/2 {
			return 4
		}
	} else if r.pos.x > r.fs.x/2 {
		if r.pos.y < r.fs.y/2 {
			return 2
		} else if r.pos.y > r.fs.y/2 {
			return 3
		}
	}

	return 0
}

func (r *robot) move(n int64) *robot {
	nx := (r.pos.x + r.vel.x*n) % r.fs.x
	ny := (r.pos.y + r.vel.y*n) % r.fs.y

	if nx < 0 {
		nx = r.fs.x + nx
	}
	if ny < 0 {
		ny = r.fs.y + ny
	}
	r.pos.x = nx
	r.pos.y = ny

	return r
}

func newRobot(fs coor, in string) robot {
	p, v, _ := strings.Cut(in, " ")
	_, p, _ = strings.Cut(p, "=")
	_, v, _ = strings.Cut(v, "=")

	return robot{
		fs:      fs,
		initPos: parsePair(p),
		pos:     parsePair(p),
		vel:     parsePair(v),
	}
}

func parsePair(in string) coor {
	x, _ := strconv.ParseInt(strings.Split(in, ",")[0], 10, 64)
	y, _ := strconv.ParseInt(strings.Split(in, ",")[1], 10, 64)
	return coor{x: x, y: y}
}
