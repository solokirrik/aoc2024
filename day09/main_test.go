package main

import (
	"context"
	_ "embed"
	"testing"
	"time"

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

func TestPart2Custom(t *testing.T) {
	for _, tCase := range []struct {
		name   string
		blocks []int64
		want   string
	}{
		{
			name:   "init",
			blocks: []int64{6, -1, -1, 9, 9, -1, -1, -1, 7, 5, 1, 8, 7},
			want:   "67199857.....",
		},
		{
			name:   "find last 7 len 1",
			blocks: []int64{6, -1, -1, 9, 9, -1, -1, -1, 7, 5, 1, 8, 7},
			want:   "67199857.....",
		},
		{
			name:   "find last 7 len 2",
			blocks: []int64{6, -1, -1, 9, 9, -1, -1, -1, 7, 5, 1, 7, 7},
			want:   "67799157.....",
		},
		{
			name:   "skip -1, find last 11 len 2",
			blocks: []int64{6, -1, -1, 9, 9, -1, -1, -1, 7, 1, 1, -1, -1},
			want:   "611997.......",
		},
		move by 1
		{
			name:   "skip -1, find 11 next 777 len 2",
			blocks: []int64{6, -1, 9, 9, -1, -1, -1, 1, 1, 7, 7, 7, -1, -1},
			want:   "69977711......",
		},
		{
			name:   "skip -1, fin 1 of 7",
			blocks: []int64{6, -1, 9, 9, -1, -1, -1, 1, 1, 7, 7, -1, 7, -1, -1},
			want:   "67997711.......",
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			s := solver{blocks: tCase.blocks}
			s.compactP2(ctx)

			assert.Equal(t, tCase.want, stringify(s.blocks))
		})
	}
}

func TestFindCompatible(t *testing.T) {
	t.Run("find last 7 len 1", func(t *testing.T) {
		s := solver{blocks: []int64{6, -1, -1, 9, 9, -1, -1, -1, 7, 5, 1, 8, 7}}
		comp := s.findFile(12, 2)
		assert.Equal(t, fileSpace{12, 1, 7}, comp)

	})

	t.Run("find last 7 len 2", func(t *testing.T) {
		s := solver{blocks: []int64{6, -1, -1, 9, 9, -1, -1, -1, 7, 5, 1, 7, 7}}
		comp := s.findFile(12, 2)
		assert.Equal(t, fileSpace{12, 2, 7}, comp)

	})

	t.Run("skip -1, find last 11 len 2", func(t *testing.T) {
		s := solver{blocks: []int64{6, -1, -1, 9, 9, -1, -1, -1, 7, 1, 1, -1, -1}}
		comp := s.findFile(len(s.blocks)-1, 2)
		assert.Equal(t, fileSpace{10, 2, 1}, comp)
	})
	t.Run("skip -1, find 1 1 next 777 len 2", func(t *testing.T) {
		s := solver{blocks: []int64{6, -1, 9, 9, -1, -1, -1, 1, 1, 7, 7, 7, -1, -1}}
		comp := s.findFile(11, 2)
		assert.Equal(t, fileSpace{8, 2, 1}, comp)

	})
	t.Run("skip -1, fin 1 of 7", func(t *testing.T) {
		s := solver{blocks: []int64{6, -1, 9, 9, -1, -1, -1, 1, 1, 7, 7, -1, 7, -1, -1}}
		comp := s.findFile(14, 1)
		assert.Equal(t, fileSpace{12, 1, 7}, comp)
	})

}
