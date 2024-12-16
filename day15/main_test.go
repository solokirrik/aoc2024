package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed ex_small
var testExS string

//go:embed ex2
var testEx2 string

//go:embed inp
var testInp string

func Test_Part1(t *testing.T) {
	t.Parallel()

	doDraw := false

	t.Run("exS", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 2028, new(solver).prep(testExS).part1(doDraw))
	})

	t.Run("ex", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 10092, new(solver).prep(testEx).part1(doDraw))
	})

	t.Run("1", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, 1413675, new(solver).prep(testInp).part1(doDraw))
	})
}

func Test_Part2(t *testing.T) {
	t.Run("ex2", func(t *testing.T) {
		assert.Equal(t, 0, new(solver).prep(testEx2).
			scaleX2().
			part2(true))
	})

	t.Run("custom-0", func(t *testing.T) {
		inp := `######
#....#
#....#
#.O..#
#.O..#
#.@..#
######

>^^`
		assert.Equal(t, 0, new(solver).prep(inp).
			scaleX2().
			part2(true))
	})

	t.Run("custom-1-vert-part", func(t *testing.T) {
		inp := `####################
##[]..[]......[][]##
##[]........@..[].##
##...........[].[]##
##............[][]##
##..##[]..[][]....##
##.....[]..[].[][]##
##........[]..[]..##
####################

>vv`
		s := new(solver).prep(inp)
		assert.Equal(t, 0, s.part2(true))
	})

	t.Run("custom-2", func(t *testing.T) {
		inp := `####################
##[]..[]......[][]##
##...........#.[].##
##..##[]..[][]....##
##[]........@..[].##
##...........[].[]##
##.....[]..[].[][]##
##........[]..[]..##
####################

^^<>>^`
		s := new(solver).prep(inp)
		assert.Equal(t, 0, s.part2(true))
	})

	t.Run("ex", func(t *testing.T) {
		s := new(solver).prep(testEx).scaleX2()
		assert.Equal(t, 9021, s.part2(false))
	})

	t.Run("2", func(t *testing.T) {
		s := new(solver).prep(testInp).scaleX2()
		assert.NotEqual(t, 1428843, s.part2(true))
	})
}

func TestScaleX2(t *testing.T) {
	inp := "#######\n" +
		"#...#.#\n" +
		"#.....#\n" +
		"#..OO@#\n" +
		"#..O..#\n" +
		"#.....#\n" +
		"#######"
	want := "##############" +
		"##......##..##" +
		"##..........##" +
		"##....[][]@.##" +
		"##....[]....##" +
		"##..........##" +
		"##############"
	got := new(solver).prep(inp).scaleX2()
	assert.Equal(t, want, fold(got.mtx))
}

func fold(mtx [][]string) string {
	var out string

	for _, row := range mtx {
		out += strings.Join(row, "")
	}

	return out
}
