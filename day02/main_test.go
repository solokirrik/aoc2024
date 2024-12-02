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
	assert.Equal(t, 2, part1(strings.Split(testEx, "\n")))
	assert.Equal(t, 402, part1(strings.Split(testInp, "\n")))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 4, part2(strings.Split(testEx, "\n")))
	assert.Equal(t, 455, part2(strings.Split(testInp, "\n")))
}
