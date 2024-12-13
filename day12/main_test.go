package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed ex1
var testEx1 string

//go:embed ex2
var testEx2 string

//go:embed ex3
var testEx3 string

//go:embed ex4
var testEx4 string

//go:embed inp
var testInp string

func Test_Part1(t *testing.T) {
	t.Run("ex1", func(t *testing.T) {
		assert.Equal(t, 4*10+4*8+4*10+4*1+3*8, new(solver).prep(testEx1).part1())
	})

	t.Run("ex2", func(t *testing.T) {
		assert.Equal(t, 21*36+1*4*4, new(solver).prep(testEx2).part1())
	})

	t.Run("ex", func(t *testing.T) {
		assert.Equal(t, 1930, new(solver).prep(testEx).part1())
	})

	t.Run("1", func(t *testing.T) {
		assert.Equal(t, 1489582, new(solver).prep(testInp).part1())
	})
}

func Test_Part2(t *testing.T) {
	t.Run("ex1", func(t *testing.T) {
		assert.Equal(t, 16+16+32+4+12, new(solver).prep(testEx1).part2())
	})

	t.Run("ex2", func(t *testing.T) {
		assert.Equal(t, 436, new(solver).prep(testEx2).part2())
	})

	t.Run("ex3", func(t *testing.T) {
		assert.Equal(t, 236, new(solver).prep(testEx3).part2())
	})

	t.Run("ex4", func(t *testing.T) {
		assert.Equal(t, 368, new(solver).prep(testEx4).part2())
	})

	t.Run("ex", func(t *testing.T) {
		assert.Equal(t, 1206, new(solver).prep(testEx).part2())
	})

	t.Run("2", func(t *testing.T) {
		assert.Equal(t, 914966, new(solver).prep(testInp).part2())
	})
}
