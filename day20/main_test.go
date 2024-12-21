package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/solokirrik/aoc2024/pkg"
	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed inp
var testInp string

func Test_Part1Ex(t *testing.T) {
	assert.Equal(t, 44, new(solver).prep(testEx).part1(1))
}

func Test_Part1(t *testing.T) {
	assert.Equal(t, 1384, new(solver).prep(testInp).part1(100))
}

func Test_Part2Ex(t *testing.T) {
	assert.Equal(t, 285, new(solver).prep(testEx).part2(50))
}

func Test_Part2(t *testing.T) {
	got := new(solver).prep(testInp).part2(100)

	assert.Equal(t, 0, got)
	assert.Greater(t, got, 946802)
	assert.NotEqual(t, got, 1031233)
	assert.Less(t, got, 1057082)
}

func Test(t *testing.T) {
	fmt.Println(manhattanDistance(pkg.NewCoord(3, 1), pkg.NewCoord(7, 3)))
	fmt.Println(manhattanDistance(pkg.NewCoord(3, 1), pkg.NewCoord(7, 3)))
}
