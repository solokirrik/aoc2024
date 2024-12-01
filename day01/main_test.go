package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testxample string

func TestP1(t *testing.T) {
	wantDiff := 11
	gotDiff := part1(testxample)
	assert.Equal(t, wantDiff, gotDiff)
}

func TestP2(t *testing.T) {
	wantDiff := 31
	gotDiff := part2(testxample)
	assert.Equal(t, wantDiff, gotDiff)
}
