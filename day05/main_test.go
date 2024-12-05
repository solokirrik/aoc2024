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
	assert.Equal(t, 143, new(solver).prep(testEx).part1())
	assert.Equal(t, 5208, new(solver).prep(testInp).part1())
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 123, new(solver).prep(testEx).part2())
	assert.Equal(t, 6732, new(solver).prep(testInp).part2())
}
