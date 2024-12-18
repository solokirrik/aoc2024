package main

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/solokirrik/aoc2024/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed ex
var testEx string

//go:embed inp
var testInp string

func Test_Part1(t *testing.T) {
	assert.Equal(t, 1928, new(solver).prep(testEx).part1())
	assert.Equal(t, 6385338159127, new(solver).prep(testInp).part1())
}

func Test_Part2_ex(t *testing.T) {
	/*
		00...111...2...333.44.5555.6666.777.888899
		0099.111...2...333.44.5555.6666.777.8888.. <- 99
		0099.1117772...333.44.5555.6666.....8888.. <- 777
		0099.111777244.333....5555.6666.....8888.. <- 44
		00992111777.44.333....5555.6666.....8888.. <- 2
	*/
	assert.Equal(t, 2858, new(solver).prep(testEx).part2())
}

func Test_Part2(t *testing.T) {
	require.Equal(t, 6415163624282, new(solver).prep(testInp).part2())
}

func TestLastFile(t *testing.T) {
	t.Parallel()

	t.Run("find last 7 len 1", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks("6..99...75187")}
		got := s.lastFile(12)
		assert.Equal(t, fileSpace{12, 1, 7}, got)
	})

	t.Run("find last 7 len 2", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks("6..99...75177")}
		got := s.lastFile(7)
		assert.Equal(t, fileSpace{4, 2, 9}, got)
	})

	t.Run("skip -1, find last 11 len 2", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks("6..99...711..")}
		got := s.lastFile(12)
		assert.Equal(t, fileSpace{10, 2, 1}, got)
	})

	t.Run("skip -1, find 1 1 next 777 len 2", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks("6.99...11777..")}
		got := s.lastFile(12)
		assert.Equal(t, fileSpace{11, 3, 7}, got)
	})

	t.Run("skip -1, fin 1 of 7", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks("6.99...1177.7..")}
		got := s.lastFile(14)
		assert.Equal(t, fileSpace{12, 1, 7}, got)
	})
}

func TestFindEmpty(t *testing.T) {
	t.Parallel()

	inp := "6.99...75177...." // len 15

	t.Run("1", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks(inp)}
		got := s.findEmpty(1, 15)
		assert.Equal(t, fileSpace{1, 1, -1}, got)
	})

	t.Run("2", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks(inp)}
		got := s.findEmpty(2, 15)
		assert.Equal(t, fileSpace{4, 3, -1}, got)
	})

	t.Run("3", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks(inp)}
		got := s.findEmpty(3, 15)
		assert.Equal(t, fileSpace{4, 3, -1}, got)
	})

	t.Run("no 4", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks(inp)}
		got := s.findEmpty(4, 11)
		assert.Equal(t, fileSpace{0, 0, -1}, got)
	})

	t.Run("4", func(t *testing.T) {
		t.Parallel()

		s := solver{blocks: stringToBlocks(inp)}
		got := s.findEmpty(4, 15)
		assert.Equal(t, fileSpace{12, 4, -1}, got)
	})
}

func stringToBlocks(in string) []int64 {
	tmp := strings.Split(in, "")
	tmp = lo.ReplaceAll(tmp, ".", "-1")

	out := make([]int64, 0, len(tmp))
	for i := range len(tmp) {
		val, err := strconv.ParseInt(tmp[i], 10, 64)
		pkg.PanicOnErr(err)
		out = append(out, val)
	}

	return out
}
