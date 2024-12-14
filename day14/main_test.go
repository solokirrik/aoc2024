package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed inp
var testInp string

func Test_Part1(t *testing.T) {
	t.Run("ex0", func(t *testing.T) {
		s := solver{robots: []robot{{pos: coor{x: 2, y: 4}, vel: coor{x: 2, y: -3}}}}
		r := s.robots[0].move(5)
		assert.Equal(t, robot{pos: coor{x: 1, y: 3}, vel: coor{x: 2, y: -3}}, r)
	})

	t.Run("ex1", func(t *testing.T) {
		s := solver{robots: []robot{{pos: coor{x: 2, y: 4}, vel: coor{x: 2, y: -3}}}}
		r := s.robots[0].move(2)
		assert.Equal(t, robot{pos: coor{x: 1, y: 3}, vel: coor{x: 2, y: -3}}, r)
	})

	t.Run("ex", func(t *testing.T) {
		assert.Equal(t, 12, new(solver).prep(testEx, 10, 7).part1(100))
	})
	t.Run("1", func(t *testing.T) {
		assert.Equal(t, 0, new(solver).prep(testInp, 101, 103).part1(100))
	})
}

func Test_Part2(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		assert.Equal(t, 0, new(solver).prep(testEx, 10, 7).part2(100))
	})
	t.Run("2", func(t *testing.T) {
		assert.Equal(t, 0, new(solver).prep(testInp, 101, 103).part2(100))
	})
}
