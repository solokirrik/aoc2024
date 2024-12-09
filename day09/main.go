package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Ans:", new(solver).prep(inp).part1())
	slog.Info("Part 2:", "Ans:", new(solver).prep(inp).part2())
}

type solver struct {
	file   []byte
	blocks []int64
}

func (s *solver) prep(inp string) *solver {
	s.file = []byte(inp)
	return s
}

const EMPTY int64 = -1

func (s *solver) part1() int {
	return s.parseBlocks().compactP1().checkSum()
}

func (s *solver) part2() int {
	return s.parseBlocks().compactP2().checkSum()
}

func (s *solver) parseBlocks() *solver {
	totalBlocks := 0
	for i := 0; i < len(s.file); i++ {
		posVal := s.file[i] - '0'
		totalBlocks += int(posVal)
	}

	b := 0
	s.blocks = make([]int64, totalBlocks)

	for i := 0; i < len(s.file); i++ {
		blockVal := EMPTY
		countBlocks := s.file[i] - '0'
		if i%2 == 0.0 {
			blockVal = int64(i / 2)
		}

		for j := byte(0); j < countBlocks; j++ {
			s.blocks[b] = blockVal
			b++
		}
	}

	s.file = s.file[:]

	return s
}

func (s *solver) compactP1() *solver {
	i, j := 0, len(s.blocks)-1

	for i < j && i < len(s.blocks) && j >= 0 {
		if s.blocks[i] != EMPTY {
			i++
			continue
		}
		if s.blocks[j] == EMPTY {
			j--
			continue
		}
		s.blocks[i], s.blocks[j] = s.blocks[j], s.blocks[i]
		i++
		j--
	}

	return s
}

type fileSpace struct {
	start, length int
	val           int64
}

func (s *solver) compactP2() *solver {
	i, j := 0, len(s.blocks)-1

	for i < j && i < len(s.blocks) && j >= 0 {
		empty := s.findEmpty(i)
		file := s.findFile(j, empty.length)

		s.swap(empty, file)

		i = empty.start + empty.length
		j = file.start - file.length
	}

	return s
}

func (s *solver) findEmpty(from int) fileSpace {
	empty := fileSpace{
		start:  0,
		length: 0,
		val:    EMPTY,
	}

	for i := from; i < len(s.blocks); i++ {
		if s.blocks[i] == EMPTY {
			if empty.length == 0 {
				empty.start = i
				empty.val = s.blocks[i]
			}
			empty.length++
		}
		if s.blocks[i] != EMPTY && empty.length > 0 {
			return empty
		}
	}

	return empty
}

func (s *solver) findFile(until int, wantLen int) fileSpace {
	i := until
	out := fileSpace{
		start:  i,
		length: 0,
		val:    s.blocks[i],
	}

	wasFirst := false
	for i > 0 {
		curChar := s.blocks[i]
		nextChar := s.blocks[i-1]

		if curChar == EMPTY {
			i--
			continue
		}

		if curChar != EMPTY {
			if out.length == 0 {
				out.start = i
				out.length = 1
				out.val = curChar
				wasFirst = true
			} else {
				out.length++
			}

			if curChar != nextChar {
				if !wasFirst {
					out.length++
				}
				if out.length <= wantLen {
					return out
				} else {
					out = fileSpace{
						start:  0,
						length: 0,
						val:    EMPTY,
					}
				}
			}
		}
		i--
	}

	return out
}

func (s *solver) swap(empty, candidate fileSpace) {
	for i := 0; i < min(empty.length, candidate.length); i++ {
		s.blocks[empty.start+i], s.blocks[candidate.start-i] = s.blocks[candidate.start-i], s.blocks[empty.start+i]
	}
}

func (s *solver) checkSum() int {
	checkSum := 0

	for i := 0; i < len(s.blocks); i++ {
		if s.blocks[i] == EMPTY {
			continue
		}
		checkSum += i * int(s.blocks[i])
	}

	return checkSum
}

func stringify(in []int64) string {
	dst := make([]string, len(in))

	for i := range dst {
		if in[i] == EMPTY {
			dst[i] = "."
		} else {
			dst[i] = strconv.FormatInt(in[i], 10)
		}
	}
	return strings.Join(dst, "")
}

func print(in []int64) {
	fmt.Println(stringify(in))
}

func save(in []int64) {
	f, err := os.Create("./p2.txt")
	panicOnErr(err)

	defer f.Close()

	f.WriteString(stringify(in) + "\n")
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
