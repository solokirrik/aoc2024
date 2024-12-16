package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex1
var testEx1 string

//go:embed ex2
var testEx2 string

//go:embed inp
var testInp string

func TestPart1Ex1(t *testing.T) {
	assert.Equal(t, 7036, new(solver).prep(testEx1).part1())
}

func TestPart1Ex2(t *testing.T) {
	assert.Equal(t, 11048, new(solver).prep(testEx2).part1())
}

func TestPart1(t *testing.T) {
	got := new(solver).prep(testInp).part1()
	assert.Equal(t, 109496, got)
}

func TestEx1Part2(t *testing.T) {
	assert.Equal(t, 45, new(solver).prep(testEx1).part2())
}

func TestEx2Part2(t *testing.T) {
	assert.Equal(t, 64, new(solver).prep(testEx2).part2())
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 551, new(solver).prep(testInp).part2())
}
