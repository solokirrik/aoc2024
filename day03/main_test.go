package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed inp
var testInp string

func Test_Part1(t *testing.T) {
	assert.Equal(t, 2*4+5*5+11*8+8*5, part1(strings.Split(testEx, "\n")))
	assert.Equal(t, 0, part1(strings.Split(testInp, "\n")))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 0, part2(strings.Split(testEx, "\n")))
	assert.Equal(t, 0, part2(strings.Split(testInp, "\n")))
}
