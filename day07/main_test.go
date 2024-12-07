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
	assert.Equal(t, 3749, new(solver).prep(testEx).part1())
	assert.Equal(t, 1298103531759, new(solver).prep(testInp).part1())
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 11387, new(solver).prep(testEx).part2())
	assert.Equal(t, 140575048428831, new(solver).prep(testInp).part2())
}
