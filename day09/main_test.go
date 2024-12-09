package main

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/samber/lo"
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
	assert.Equal(t, 2858, new(solver).prep(testEx).part2())
}

func Test_Part2(t *testing.T) {
	check := new(solver).prep(testInp).part2()
	require.Less(t, check, 13450752761922)
	require.Less(t, check, 13127419183747)
	require.Less(t, check, 6385338159127)
	require.Greater(t, check, 6285519204396)
	t.Log(check)
}

func TestFindCompatible(t *testing.T) {
	t.Run("find last 7 len 1", func(t *testing.T) {
		s := solver{blocks: stringToBlocks("6..99...75187")}
		comp := s.findFile(12, 2)
		assert.Equal(t, fileSpace{12, 1, 7}, comp)

	})

	t.Run("find last 7 len 2", func(t *testing.T) {
		s := solver{blocks: stringToBlocks("6..99...75177")}
		comp := s.findFile(12, 2)
		assert.Equal(t, fileSpace{12, 2, 7}, comp)

	})

	t.Run("skip -1, find last 11 len 2", func(t *testing.T) {
		s := solver{blocks: stringToBlocks("6..99...711..")}
		comp := s.findFile(len(s.blocks)-1, 2)
		assert.Equal(t, fileSpace{10, 2, 1}, comp)
	})
	t.Run("skip -1, find 1 1 next 777 len 2", func(t *testing.T) {
		s := solver{blocks: stringToBlocks("6.99...11777..")}
		comp := s.findFile(11, 2)
		assert.Equal(t, fileSpace{8, 2, 1}, comp)

	})
	t.Run("skip -1, fin 1 of 7", func(t *testing.T) {
		s := solver{blocks: stringToBlocks("6.99...1177.7..")}
		comp := s.findFile(14, 1)
		assert.Equal(t, fileSpace{12, 1, 7}, comp)
	})

}
func stringToBlocks(in string) []int64 {
	tmp := strings.Split(in, "")
	tmp = lo.ReplaceAll(tmp, ".", "-1")

	out := make([]int64, 0, len(tmp))
	for i := range len(tmp) {
		val, err := strconv.ParseInt(tmp[i], 10, 64)
		panicOnErr(err)
		out = append(out, val)
	}

	return out
}
