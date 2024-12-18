package main

import (
	"strconv"
)

type step struct {
	parent  dcoord
	dpos    dcoord
	score   int
	pathLen int
}

func newCoord(r, c int) coord {
	return coord{r: r, c: c}
}

type coord struct {
	r, c int
}

func (c coord) eq(p coord) bool {
	return c.r == p.r && c.c == p.c
}

func (c coord) str() string {
	return strconv.Itoa(c.c) + "," + strconv.Itoa(c.r)
}

type coordHash [2]int

func (c *coord) hash() coordHash {
	return [2]int{c.c, c.r}
}

func newDCoord(point coord, dir int) dcoord {
	nd := dcoord{c: point, dir: dir}
	return nd
}

type dcoord struct {
	c   coord
	dir int
}

type dCoordHash [3]int

func (c *dcoord) hash() dCoordHash {
	return [3]int{c.dir, c.c.c, c.c.r}
}

func (c *dcoord) eqPos(p coord) bool {
	return c.getPos().eq(p)
}

func (c *dcoord) getPos() coord {
	return newCoord(c.c.r, c.c.c)
}
