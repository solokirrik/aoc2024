package pkg

type Stack struct {
	q   [][]Coord
	mem map[[2]Coord]int
}

func NewSt() *Stack {
	return &Stack{
		mem: make(map[[2]Coord]int),
	}
}

func (s *Stack) Len() int {
	return len(s.q)
}

func (s *Stack) Push(val []Coord) *Stack {
	key := [2]Coord{val[0], {}}
	if len(val) == 2 {
		key = [2]Coord{val[0], val[1]}
	}

	if key[0].C > key[1].C {
		key[0], key[1] = key[1], key[0]
	}

	if _, ok := s.mem[key]; ok {
		return s
	}

	s.mem[key]++
	s.q = append(s.q, val)

	return s
}

func (s *Stack) Pop() []Coord {
	lastIdx := len(s.q) - 1
	val := s.q[lastIdx]
	s.q = s.q[:lastIdx]

	return val
}
