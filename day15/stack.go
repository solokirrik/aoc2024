package main

type stack struct {
	q   [][]coord
	mem map[[2]coord]int
}

func newSt() *stack {
	return &stack{
		mem: make(map[[2]coord]int),
	}
}

func (s *stack) len() int {
	return len(s.q)
}

func (s *stack) push(val []coord) *stack {
	key := [2]coord{val[0], {}}
	if len(val) == 2 {
		key = [2]coord{val[0], val[1]}
	}

	if key[0].c > key[1].c {
		key[0], key[1] = key[1], key[0]
	}

	if _, ok := s.mem[key]; ok {
		return s
	}

	s.mem[key]++
	s.q = append(s.q, val)

	return s
}

func (s *stack) pop() []coord {
	lastIdx := len(s.q) - 1
	val := s.q[lastIdx]
	s.q = s.q[:lastIdx]

	return val
}
