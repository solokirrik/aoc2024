package pkg

import (
	"strconv"
	"strings"
)

type Direction int

const (
	NORTH Direction = 0
	EAST  Direction = 1
	SOUTH Direction = 2
	WEST  Direction = 3
)

var (
	Deltas = map[Direction][2]int{
		NORTH: {-1, 0},
		EAST:  {0, 1},
		SOUTH: {1, 0},
		WEST:  {0, -1},
	}

	Opposites = map[Direction]Direction{
		NORTH: SOUTH,
		EAST:  WEST,
		SOUTH: NORTH,
		WEST:  EAST,
	}
)

type Step struct {
	Parent  DCoord
	DPos    DCoord
	Score   int
	PathLen int
	Path    []Coord
}

func NewCoord(r, c int) Coord {
	return Coord{R: r, C: c}
}

func ParseCoord(in string) Coord {
	x, _ := strconv.ParseInt(strings.Split(in, ",")[0], 10, 64)
	y, _ := strconv.ParseInt(strings.Split(in, ",")[1], 10, 64)
	return Coord{C: int(x), R: int(y)}
}

type Coord struct {
	R, C int
}

func (c Coord) Eq(p Coord) bool {
	return c.R == p.R && c.C == p.C
}

func (c Coord) Str() string {
	return strconv.Itoa(c.C) + "," + strconv.Itoa(c.R)
}

type CoordHash [2]int

func (c *Coord) Hash() CoordHash {
	return [2]int{c.C, c.R}
}

func NewDCoord(point Coord, dir Direction) DCoord {
	nd := DCoord{C: point, Dir: dir}
	return nd
}

type DCoord struct {
	C   Coord
	Dir Direction
}

type DCoordHash [3]int

func (c *DCoord) hash() DCoordHash {
	return [3]int{int(c.Dir), c.C.C, c.C.R}
}

func (c *DCoord) EqCoord(p Coord) bool {
	return c.GetPos().Eq(p)
}

func (c *DCoord) GetPos() Coord {
	return NewCoord(c.C.R, c.C.C)
}
