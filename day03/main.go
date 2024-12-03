package main

import (
	_ "embed"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1:", "Mul:", part1(inp))
	slog.Info("Part 2:", "Mul:", part2(inp))
}

var (
	mulRex = regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	doRex  = regexp.MustCompile("do\\(\\)|don't\\(\\)")
)

func part1(row string) int {
	return calcMulsRow(row)
}

func part2(row string) int {
	indexes := lo.Map(doRex.FindAllStringIndex(row, -1), func(startEndPair []int, _ int) int { return startEndPair[0] })

	if indexes[0] != 0 {
		indexes = append([]int{0}, indexes...)
	}

	indexes = append(indexes, len(row))

	sum := 0

	for i := 1; i < len(indexes); i++ {
		l, r := indexes[i-1], indexes[i]
		substr := row[l:r]

		if strings.HasPrefix(substr, "don't()") {
			continue
		}

		sum += calcMulsRow(substr)
	}

	return sum
}

func calcMulsRow(row string) int {
	pairs := lo.Map(mulRex.FindAllStringSubmatch(row, -1), func(item []string, _ int) []int {
		vals := strings.Split(item[0][4:len(item[0])-1], ",")
		n1, _ := strconv.ParseInt(vals[0], 10, 64)
		n2, _ := strconv.ParseInt(vals[1], 10, 64)
		return []int{int(n1), int(n2)}
	})

	mults := lo.Map(pairs, func(inner []int, _ int) int {
		return lo.Reduce(inner, func(acc, item int, _ int) int { return acc * item }, 1)
	})

	sum := lo.Reduce(mults, func(acc, item int, _ int) int { return acc + item }, 0)

	return sum
}
