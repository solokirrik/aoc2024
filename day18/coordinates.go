package main

import (
	"strconv"
	"strings"
)

type step struct {
	parent  dcoord
	dpos    dcoord
	score   int
	pathLen int
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

func (c *coord) hash() string {
	return strconv.Itoa(c.r) + "-" + strconv.Itoa(c.c)
}

func pathHash(path []coord) string {
	b := strings.Builder{}
	for i := range path {
		b.WriteString(path[i].hash())
	}

	return b.String()
}

type dcoord struct {
	c   coord
	dir string
	id  string
}

func (c *dcoord) hash() string {
	return c.dir + "-" + strconv.Itoa(c.c.r) + "-" + strconv.Itoa(c.c.c)
}

func (c *dcoord) eqPos(p coord) bool {
	return c.getPos().eq(p)
}

func (c *dcoord) getPos() coord {
	return newCoord(c.c.r, c.c.c)
}

func newCoord(r, c int) coord {
	return coord{r: r, c: c}
}

func newDCoord(point coord, dir string) dcoord {
	nd := dcoord{c: point, dir: dir}
	nd.id = nd.hash()
	return nd
}
