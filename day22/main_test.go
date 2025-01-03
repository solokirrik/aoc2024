package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed inp
var testInp string

func Test_Part1Ex(t *testing.T) {
	assert.Equal(t, uint64(37327623), new(solver).prep(testEx).part1())
}

func Test_Part1(t *testing.T) {
	got := new(solver).prep(testInp).part1()
	assert.Equal(t, uint64(0), got)
	fmt.Println(got)
}

func Test_Part2Ex(t *testing.T) {
	assert.Equal(t, 0, new(solver).prep(testEx).part2())
}

func Test_Part2(t *testing.T) {
	assert.Equal(t, 0, new(solver).prep(testInp).part2())
}

func TestSecretNumber(t *testing.T) {
	for i, val := range []struct {
		in      uint64
		wantOut uint64
	}{
		{in: 123, wantOut: 15887950},
		{in: 15887950, wantOut: 16495136},
		{in: 16495136, wantOut: 527345},
		{in: 527345, wantOut: 704524},
		{in: 704524, wantOut: 1553684},
		{in: 1553684, wantOut: 12683156},
		{in: 12683156, wantOut: 11100544},
		{in: 11100544, wantOut: 12249484},
		{in: 12249484, wantOut: 7753432},
		{in: 7753432, wantOut: 5908254},
	} {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			assert.Equal(t, val.wantOut, secretNumber(val.in))
		})
	}
}

func TestPrune(t *testing.T) {
	assert.Equal(t, uint64(16113920), prune(100000000))
}

func TestMix(t *testing.T) {
	assert.Equal(t, uint64(37), mix(42, 15))
}
