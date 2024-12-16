package main

import (
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"
)

//go:embed inp
var inp string

//go:embed ex_small
var exS string

//go:embed ex
var ex string

func main() {
	// slog.Info("Part 1:", "Ans", new(solver).prep(ex).part1(true))
	slog.Info("Part 2:", "Ans:", new(solver).prep(ex).part2(true))
	// slog.Info("Part 1:", "Ans", new(solver).prep(ex).part1(true))
	slog.Info("Part 2:", "Ans:", new(solver).prep(ex).part2(true))
}

type solver struct {
	mtx     [][]string
	start   coord
	moves   []string
	q       []coord
	doPrint bool
}

type robot struct {
	start coord
}

type coord struct {
	r, c int
}

func (c *coord) isZero() bool {
	return c.c == 0 && c.r == 0
}

func (c *coord) right() coord {
	return coord{r: c.r, c: c.c + 1}
}

func (c *coord) left() coord {
	return coord{r: c.r, c: c.c - 1}
}

const (
	UP    = "^"
	RIGHT = ">"
	DOWN  = "v"
	LEFT  = "<"
)

func (s *solver) prep(inp string) *solver {
	field, moves, _ := strings.Cut(inp, "\n\n")

	s.moves = strings.Split(strings.ReplaceAll(moves, "\n", ""), "")
	rows := strings.Split(field, "\n")
	s.mtx = make([][]string, 0, len(rows))

	for i, row := range rows {
		s.mtx = append(s.mtx, strings.Split(row, ""))
		if c := strings.Index(row, "@"); c != -1 {
			s.start = coord{i, c}
		}
	}

	return s
}

func (s *solver) part1(doDraw bool) int {
	s.doPrint = doDraw
	s.print(0)

	curCell := s.start
	for _, dir := range s.moves {
		curCell = s.rootMoveP1(curCell, dir)

		if doDraw {
			time.Sleep(100 * time.Millisecond)
			clean(len(s.mtx))
			s.print(0)
		}
	}

	return calcGPS(s.mtx, "O")
}

func (s *solver) part2(doPrint bool) int {
	s.doPrint = doPrint
	s.print(0)

	curCell := s.start

	i := 0
	for _, dir := range s.moves {
		nextCell := s.getNextCell(curCell, dir)
		nextCellChar := s.getChar(nextCell)

		if nextCellChar == "." {
			s.swap(curCell, nextCell)
			curCell = nextCell
			s.print(i)
			i++
			continue
		}

		if nextCellChar == "#" {
			s.print(i)
			i++
			continue
		}

		if isHorizontal(dir) {
			if s.moveHorizontal(curCell, dir) {
				curCell = nextCell
			}
			s.print(i)
			i++
			continue
		}

		nextNeighbourCell := getPair(nextCellChar, nextCell)
		next := []coord{nextCell, nextNeighbourCell}
		st := s.moveVertical(&stack{}, next, dir)

		if st.len() == 0 {
			i++
			continue
		}

		for st.len() > 0 {
			if entry := st.pop(); len(entry) == 2 {
				s.move2(entry[0], entry[1], dir)
				s.print(i)
			}
		}

		s.swap(curCell, nextCell)
		curCell = nextCell
		s.print(i)
		i++
	}

	s.print(i)

	return calcGPS(s.mtx, "[")
}

func (s *solver) moveHorizontal(cell coord, dir string) bool {
	char := s.getChar(cell)
	switch char {
	case "#":
		return false
	case ".":
		return true
	}

	nextCell := s.getNextCell(cell, dir)
	if s.moveHorizontal(nextCell, dir) {
		s.swap(cell, nextCell)
		return true
	}

	return false
}

func (s *solver) moveVertical(st *stack, cells []coord, dir string) *stack {
	c0, c1 := cells[0], cells[1]
	c0Char, c1Char := s.getChar(c0), s.getChar(c1)
	if c0Char == "#" || c1Char == "#" {
		return &stack{}
	}
	if c0Char == "." && c1Char == "." {
		return st
	}

	next0, next0Char := s.getNextCell(c0, dir), s.getChar(s.getNextCell(c0, dir))
	next0Pair := getPair(next0Char, next0)

	next1, next1Char := s.getNextCell(c1, dir), s.getChar(s.getNextCell(c1, dir))
	next1Pair := getPair(next1Char, next1)

	st = st.push(cells)

	if next0Char == "." && next1Char == "." {
		return st
	}

	//simmetrics from 1 cube on another
	if next0Pair == next1 && !next0Pair.isZero() {
		return s.moveVertical(st, []coord{next0, next0Pair}, dir)
	}

	// half cube next
	if next0Char != "." && next1Char == "." {
		return s.moveVertical(st, []coord{next0, next0Pair}, dir)
	}

	if next0Char == "." && next1Char != "." {
		return s.moveVertical(st, []coord{next1, next1Pair}, dir)
	}

	// two cubes next
	// if next0Char != "." && next1Char != "." {
	// }

	st = s.moveVertical(st, []coord{next0, next0Pair}, dir)
	if st.len() == 0 {
		return st
	}

	return s.moveVertical(st, []coord{next1, next1Pair}, dir)
}

func (s *solver) rootMoveP1(curCell coord, dir string) coord {
	for {
		s.print(0)

		nextCell := s.getNextCell(curCell, dir)
		nextSign := s.getChar(nextCell)
		switch nextSign {
		case ".":
			s.swap(curCell, nextCell)
			return nextCell
		case "#":
			return curCell
		case "O":
			if !s.move(nextCell, dir) {
				return curCell
			}
		}
	}
}

func (s *solver) getNextCell(curCell coord, dir string) coord {
	switch dir {
	case UP:
		return coord{curCell.r - 1, curCell.c}
	case RIGHT:
		return coord{curCell.r, curCell.c + 1}
	case DOWN:
		return coord{curCell.r + 1, curCell.c}
	case LEFT:
		return coord{curCell.r, curCell.c - 1}
	}

	return coord{}
}

func (s *solver) getChar(cell coord) string {
	return s.mtx[cell.r][cell.c]
}

func (s *solver) swap(a, b coord) {
	s.mtx[a.r][a.c], s.mtx[b.r][b.c] = s.mtx[b.r][b.c], s.mtx[a.r][a.c]
}
func (s *solver) move2(a, b coord, dir string) {
	aNext, bNext := s.getNextCell(a, dir), s.getNextCell(b, dir)
	s.swap(a, aNext)
	s.swap(b, bNext)
	s.print(0)
}

func (s *solver) move(cell coord, dir string) bool {
	curSign := s.getChar(cell)
	nextCell := s.getNextCell(cell, dir)
	nextSign := s.getChar(nextCell)

	switch nextSign {
	case ".":
		if isVertical(dir) && isBrackets(curSign) {
			nextNeighbourCell := s.getNextCell(getPair(curSign, cell), dir)
			nextNeighbourSign := s.getChar(nextNeighbourCell)
			if nextNeighbourSign == "." {
				s.move2(cell, getPair(curSign, cell), dir)
				return true
			}
			if nextNeighbourSign == "#" {
				return false
			}
			return s.move(nextNeighbourCell, dir)
		} else {
			s.swap(cell, nextCell)
			s.print(0)
			return true
		}
	case "#":
		return false
	case "O":
		return s.move(nextCell, dir)
	case "[", "]":
		if isHorizontal(dir) {
			return s.move(nextCell, dir)
		}

		nextNeighbourCell := getPair(nextSign, nextCell)
		return s.move(nextCell, dir) && s.move(nextNeighbourCell, dir)
	}

	return false
}

func (s *solver) scaleX2() *solver {
	newMtx := make([][]string, 0, len(s.mtx))

	for _, row := range s.mtx {
		newRow := make([]string, 0, len(row)*2)
		for _, char := range row {
			switch char {
			case "#":
				newRow = append(newRow, "#", "#")
			case ".":
				newRow = append(newRow, ".", ".")
			case "O":
				newRow = append(newRow, "[", "]")
			case "@":
				s.start.c = len(newRow)
				newRow = append(newRow, "@", ".")
			}
		}
		newMtx = append(newMtx, newRow)
	}

	s.mtx = newMtx

	return s
}

func getPair(sign string, cell coord) coord {
	switch sign {
	case "]":
		return cell.left()
	case "[":
		return cell.right()
	}

	return coord{}
}

func swap(mtx [][]string, currCell, nextCell coord) bool {
	if mtx[nextCell.r][nextCell.c] == "." {
		mtx[currCell.r][currCell.c], mtx[nextCell.r][nextCell.c] = mtx[nextCell.r][nextCell.c], mtx[currCell.r][currCell.c]
		return true
	}

	return false
}

func isHorizontal(dir string) bool {
	return slices.Contains([]string{RIGHT, LEFT}, dir)
}

func isVertical(dir string) bool {
	return slices.Contains([]string{UP, DOWN}, dir)
}

func isBrackets(sign string) bool {
	return slices.Contains([]string{"[", "]"}, sign)
}

func calcGPS(mtx [][]string, char string) int {
	gps := 0

	for i := range mtx {
		for j := range mtx[i] {
			if mtx[i][j] == char {
				gps += i*100 + j
			}
		}
	}

	return gps
}

func (s *solver) print(i int) {
	if !s.doPrint {
		return
	}

	var out string
	fmt.Println(i)
	for _, row := range s.mtx {
		out += strings.Join(row, "") + "\n"
	}

	fmt.Println(out)
}

type stack struct {
	q [][]coord
}

func (s *stack) len() int {
	return len(s.q)
}

func (s *stack) push(val []coord) *stack {
	s.q = append(s.q, val)

	return s
}

func (s *stack) pop() []coord {
	lastIdx := len(s.q) - 1
	val := s.q[lastIdx]
	s.q = s.q[:lastIdx]

	return val
}

const ESC = 27

var clear = fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)

func clean(linesY int) {
	_, _ = fmt.Fprint(io.Writer(os.Stdout), strings.Repeat(clear, linesY+2))
}
