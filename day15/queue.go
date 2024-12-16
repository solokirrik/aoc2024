package main

func newQQ() *queue {
	return &queue{
		mem: make(map[[2]coord]int),
	}
}

type queue struct {
	q   [][]coord
	mem map[[2]coord]int
}

func (q *queue) len() int {
	return len(q.q)
}

func (q *queue) push(val []coord) *queue {
	key := [2]coord{val[0], {}}
	if len(val) == 2 {
		key = [2]coord{val[0], val[0]}
	}

	if key[0].c > key[1].c {
		key[0], key[1] = key[1], key[0]
	}

	if _, ok := q.mem[key]; ok {
		return q
	}

	q.q = append(q.q, val)
	q.mem[key]++
	return q
}

func (q *queue) get() []coord {
	val := q.q[0]
	q.q = q.q[1:]

	return val
}
