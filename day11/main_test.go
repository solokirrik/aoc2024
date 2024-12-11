package main

import (
	_ "embed"
	"strconv"
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
	t.Run("ex-1", func(t *testing.T) {
		assert.Equal(t, 7, new(solver).prep(testEx).part1(1))
	})

	t.Run("ex2-6", func(t *testing.T) {
		assert.Equal(t, 22, new(solver).prep(testEx2).part1(6))
	})

	t.Run("ex2-25", func(t *testing.T) {
		assert.Equal(t, 55312, new(solver).prep(testEx2).part1(25))
	})

	t.Run("1-25", func(t *testing.T) {
		assert.Equal(t, 193607, new(solver).prep(testInp).part1(25))
	})
}

func Test_Part2(t *testing.T) {
	t.Run("ex-1", func(t *testing.T) {
		assert.Equal(t, 7, new(solver).prep(testEx).part2(1))
	})

	t.Run("ex2-6", func(t *testing.T) {
		assert.Equal(t, 22, new(solver).prep(testEx2).part2(6))
	})

	t.Run("ex2-25", func(t *testing.T) {
		assert.Equal(t, 55312, new(solver).prep(testEx2).part2(25))
	})

	t.Run("2-25", func(t *testing.T) {
		assert.Equal(t, 193607, new(solver).prep(testInp).part2(25))
	})

	t.Run("2-75", func(t *testing.T) {
		assert.Equal(t, 229557103025807, new(solver).prep(testInp).part2(75))
	})
}

func TestRules(t *testing.T) {
	for _, tCase := range []struct {
		n    int64
		want []int64
	}{
		{
			n:    0,
			want: []int64{1},
		},
		{
			n:    1,
			want: []int64{2024},
		},
		{
			n:    10,
			want: []int64{1, 0},
		},
		{
			n:    99,
			want: []int64{9, 9},
		},
		{
			n:    999,
			want: []int64{2021976},
		},
	} {
		t.Run(strconv.FormatInt(tCase.n, 10), func(t *testing.T) {
			t.Parallel()
			got := applyRule(tCase.n)
			assert.ElementsMatch(t, tCase.want, got)
		})
	}
}
