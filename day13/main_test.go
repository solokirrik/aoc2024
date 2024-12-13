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
	t.Run("ex", func(t *testing.T) {
		assert.Equal(t, uint64(480), new(solver).prep(testEx, 0).part1())
	})

	t.Run("inp", func(t *testing.T) {
		assert.Equal(t, uint64(32041), new(solver).prep(testInp, 0).part1())
	})
}

func Test_Part2(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		assert.Equal(t, uint64(875318608908), new(solver).prep(testEx, p2PrizeOffset).part2())
	})

	t.Run("inp", func(t *testing.T) {
		assert.Equal(t, uint64(95843948914827), new(solver).prep(testInp, p2PrizeOffset).part2())
	})
}

func TestSolve(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		wantN, wantK := float64(80), float64(40)
		gotN, gotK := solve(button{94, 34}, button{22, 67}, button{8400, 5400})
		assert.Equal(t, wantN, gotN)
		assert.Equal(t, wantK, gotK)
	})

	t.Run("2-no", func(t *testing.T) {
		wantN, wantK := float64(0), float64(0)
		gotN, gotK := solve(button{26, 66}, button{67, 21}, button{12748, 12176})
		assert.Equal(t, wantN, gotN)
		assert.Equal(t, wantK, gotK)
	})

	t.Run("3", func(t *testing.T) {
		wantN, wantK := float64(38), float64(86)
		gotN, gotK := solve(button{17, 86}, button{84, 37}, button{7870, 6450})
		assert.Equal(t, wantN, gotN)
		assert.Equal(t, wantK, gotK)
	})
}
