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
	assert.Equal(t, 0, part1(testEx))
	assert.Equal(t, 0, part1(testInp))
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 0, part2(testEx))
	assert.Equal(t, 0, part2(testInp))
}
