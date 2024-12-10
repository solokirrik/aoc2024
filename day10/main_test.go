package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex4
var testEx4 string

//go:embed inp
var testInp string

func Test_Part1(t *testing.T) {
	t.Run("ex4", func(t *testing.T) {
		assert.Equal(t, 36, new(solver).prep(testEx4).part1())
	})

	t.Run("1", func(t *testing.T) {
		assert.Equal(t, 698, new(solver).prep(testInp).part1())
	})
}

func Test_Part2(t *testing.T) {
	t.Run("ex4", func(t *testing.T) {
		assert.Equal(t, 81, new(solver).prep(testEx4).part2())
	})

	t.Run("1", func(t *testing.T) {
		assert.Equal(t, 1436, new(solver).prep(testInp).part2())
	})
}
