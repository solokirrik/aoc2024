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

const ()

func Test_Part1Ex(t *testing.T) {
	got := new(solver).prep(maxGridEx, bytesEx, testEx).part1(maxGridEx)
	assert.Equal(t, 22, got)
}

func Test_Part1(t *testing.T) {
	got := new(solver).prep(maxGrid, bytesN, testInp).part1(maxGrid)
	assert.Equal(t, 322, got)
}

func Test_Part2Ex(t *testing.T) {
	got := part2(maxGridEx, bytesEx, testEx)
	assert.Equal(t, "6,1", got)
}

func Test_Part2(t *testing.T) {
	got := part2(maxGrid, bytesN, inp)
	assert.Equal(t, "60,21", got)
}
