package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed inp
var testInp string

func Test_Part1(t *testing.T) {
	t.Run("quadrant", func(t *testing.T) {
		fs := coor{11, 7}

		for _, tCase := range []struct {
			rx, ry int64
			wantQ  int
		}{
			{rx: 1, ry: 1, wantQ: 1},
			{rx: 2, ry: 2, wantQ: 1},
			{rx: 4, ry: 1, wantQ: 1},

			{rx: 9, ry: 1, wantQ: 2},
			{rx: 9, ry: 5, wantQ: 3},
			{rx: 1, ry: 9, wantQ: 4},

			{rx: 5, ry: 0, wantQ: 0},
			{rx: 5, ry: 1, wantQ: 0},
			{rx: 5, ry: 2, wantQ: 0},
			{rx: 5, ry: 3, wantQ: 0},
			{rx: 5, ry: 4, wantQ: 0},
			{rx: 5, ry: 5, wantQ: 0},
			{rx: 5, ry: 6, wantQ: 0},

			{rx: 0, ry: 3, wantQ: 0},
			{rx: 1, ry: 3, wantQ: 0},
			{rx: 2, ry: 3, wantQ: 0},
			{rx: 3, ry: 3, wantQ: 0},
			{rx: 4, ry: 3, wantQ: 0},
			{rx: 5, ry: 3, wantQ: 0},
			{rx: 6, ry: 3, wantQ: 0},
			{rx: 7, ry: 3, wantQ: 0},
			{rx: 8, ry: 3, wantQ: 0},
			{rx: 9, ry: 3, wantQ: 0},
			{rx: 10, ry: 3, wantQ: 0},
		} {
			t.Run(fmt.Sprintf("%d:%d/%d", tCase.rx, tCase.ry, tCase.wantQ), func(t *testing.T) {
				t.Parallel()

				r := robot{fs: fs, pos: coor{tCase.rx, tCase.ry}}
				assert.Equal(t, tCase.wantQ, r.q())
			})
		}
	})

	t.Run("new-robor", func(t *testing.T) {
		fs := coor{11, 7}

		for _, tCase := range []struct {
			inp     string
			wantPos coor
			wantVel coor
		}{
			{
				inp:     "p=0,4 v=3,-3",
				wantPos: coor{0, 4},
				wantVel: coor{3, -3},
			},
			{
				inp:     "p=6,3 v=-1,-3",
				wantPos: coor{6, 3},
				wantVel: coor{-1, -3},
			},
			{
				inp:     "p=10,3 v=-1,2",
				wantPos: coor{10, 3},
				wantVel: coor{-1, 2},
			},
		} {
			t.Run(tCase.inp, func(t *testing.T) {
				t.Parallel()

				rob := newRobot(fs, tCase.inp)
				assert.Equal(t, tCase.wantPos, rob.pos)
				assert.Equal(t, tCase.wantPos, rob.initPos)
				assert.Equal(t, tCase.wantVel, rob.vel)
			})
		}
	})

	t.Run("ex1-2", func(t *testing.T) {
		fs := coor{11, 7}

		for _, tCase := range []struct {
			r       robot
			move    int64
			wantPos coor
			wantQ   int
		}{
			{
				r:       robot{fs: fs, pos: coor{x: 2, y: 4}, vel: coor{x: 2, y: -3}},
				move:    2,
				wantPos: coor{6, 5},
				wantQ:   3,
			},
		} {
			t.Run(fmt.Sprintf("%v", tCase.r.pos), func(t *testing.T) {
				tCase.r.move(tCase.move)
				assert.Equal(t, tCase.wantPos, tCase.r.pos)
				assert.Equal(t, tCase.wantQ, tCase.r.q())
			})
		}
	})

	t.Run("ex", func(t *testing.T) {
		fs := coor{11, 7}

		got := new(solver).prep(testEx, fs.x, fs.y).part1(100)
		assert.Equal(t, 12, got)
	})

	t.Run("1", func(t *testing.T) {
		assert.Equal(t, 0, new(solver).prep(testInp, 101, 103).part1(100))
	})
}

func Test_Part2(t *testing.T) {
	t.Run("2", func(t *testing.T) {
		new(solver).prep(testInp, fieldW, fieldH).part2(10000, false)
	})

	t.Run("draw", func(t *testing.T) {
		new(solver).prep(testInp, fieldW, fieldH).part2n(7520)
	})
}
