package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed ex2
var testEx2 string

//go:embed inp
var testInp string

func Test_Part1(t *testing.T) {
	assert.Equal(t, 2*4+5*5+11*8+8*5, part1(testEx))
	assert.Equal(t, 192767529, part1(testInp))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 2*4+8*5, part2(testEx2))
	assert.Equal(t, 104083373, part2(testInp))
}
