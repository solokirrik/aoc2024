package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed ex
var testEx string

//go:embed ex2
var testEx2 string

//go:embed inp
var testInp string

func Test_SmallEx(t *testing.T) {
	t.Parallel()

	type state struct {
		a   uint64
		b   uint64
		c   uint64
		ops []uint64
	}
	type want struct {
		a   uint64
		b   uint64
		c   uint64
		out string
	}

	for _, tCase := range []struct {
		name string
		got  state
		want want
	}{
		{
			name: "1",
			got: state{
				c:   9,
				ops: []uint64{2, 6},
			},
			want: want{
				b:   1,
				c:   9,
				out: "",
			},
		},
		{
			name: "2",
			got: state{
				a:   10,
				ops: []uint64{5, 0, 5, 1, 5, 4},
			},
			want: want{
				a:   10,
				out: "0,1,2",
			},
		},
		{
			name: "3",
			got: state{
				a:   2024,
				ops: []uint64{0, 1, 5, 4, 3, 0},
			},
			want: want{
				a:   0,
				out: "4,2,5,6,7,7,7,7,3,1,0",
			},
		},
		{
			name: "4",
			got: state{
				b:   29,
				ops: []uint64{1, 7},
			},
			want: want{b: 26},
		},
		{
			name: "5",
			got: state{
				b:   2024,
				c:   43690,
				ops: []uint64{4, 0},
			},
			want: want{
				b:   44354,
				c:   43690,
				out: "",
			},
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			t.Parallel()

			s := solver{
				regA:  tCase.got.a,
				regB:  tCase.got.b,
				regC:  tCase.got.c,
				instr: tCase.got.ops,
			}
			got := s.part1()
			assert.Equal(t, tCase.want.a, s.regA, "A")
			assert.Equal(t, tCase.want.b, s.regB, "B")
			assert.Equal(t, tCase.want.c, s.regC, "C")
			assert.Equal(t, tCase.want.out, got, "out")
		})
	}
}

func Test_Part1Ex(t *testing.T) {
	got := new(solver).prep(testEx).part1()
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", got)
}

func Test_Part1Inp(t *testing.T) {
	got := new(solver).prep(testInp).part1()
	assert.Equal(t, "2,1,3,0,5,2,3,7,1", got)
}

func Test_Part2Ex(t *testing.T) {
	got := new(solver).prep(testEx2).part2()
	assert.Equal(t, 117440, got)
}

func Test_Part2Inp(t *testing.T) {
	assert.Equal(t, "", new(solver).prep(testInp).part2())
}
