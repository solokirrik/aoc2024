package main

import "slices"

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
