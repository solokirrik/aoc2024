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

func Test_Part1Ex(t *testing.T) {
	got := new(solver).prep(testEx).part1()
	assert.Equal(t, 68*29+60*980+68*179+64*456+64*379, got)
}

func Test_Part1(t *testing.T) {
	got := new(solver).prep(testInp).part1()
	assert.Equal(t, 0, got)
}

func Test_Part2Ex(t *testing.T) {
	assert.Equal(t, 0, new(solver).prep(testEx).part2())
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 0, new(solver).prep(testInp).part2())
}
