package main

import (
	"slices"
	"strconv"
	"strings"
)

type step struct {
	parent dcoord
	pos    dcoord
	score  int
	path   []coord
}

type coord struct {
	r, c int
}

func (c coord) eq(p coord) bool {
	return c.r == p.r && c.c == p.c
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

func newQQ() queue {
	return queue{q: make([]step, 0, 50)}
}

type queue struct {
	q []step
}

func (q *queue) push(s step) *queue {
	q.q = append(q.q, s)
	return q
}

func (q *queue) sort() {
	slices.SortFunc(q.q, func(a, b step) int {
		return a.score - b.score
	})
}

func (q *queue) len() int {
	return len(q.q)
}

func (q *queue) get() step {
	val := q.q[0]
	q.q = q.q[1:]
	return val
}

func newCoord(r, c int) coord {
	return coord{r: r, c: c}
}

func newDCoord(point coord, dir string) dcoord {
	nd := dcoord{c: point, dir: dir}
	nd.id = nd.hash()
	return nd
}
