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
	want := 7
	assert.Equal(t, want, got)
}

func Test_Part1(t *testing.T) {
	got := new(solver).prep(testInp).part1()
	want := 1366
	assert.Equal(t, want, got)
}

func Test_Part2Ex(t *testing.T) {
	got := new(solver).prep(testEx).part2()
	want := "co,de,ka,ta"
	assert.Equal(t, want, got)
}

func Test_Part2(t *testing.T) {
	got := new(solver).prep(testInp).part2()
	want := "bs,cf,cn,gb,gk,jf,mp,qk,qo,st,ti,uc,xw"
	assert.Equal(t, want, got)
}
