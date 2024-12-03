package main

import (
	_ "embed"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed ex
var inp string

func main() {
	inplines := strings.Split(inp, "\n")

	slog.Info("Part 1:", "safe:", part1(inplines))
	slog.Info("Part 2:", "safe:", part2(inplines))
}

func part1(inp []string) int64 {
	result := int64(0)
	for _, row := range inp {
		result += calcMulsRow(row)
	}

	return result
}

var rego = regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")

func calcMulsRow(row string) int64 {
	pairs := lo.Map(rego.FindAllStringSubmatch(row, -1), func(item []string, _ int) []int64 {
		vals := strings.Split(item[0][4:len(item[0])-1], ",")
		n1, _ := strconv.ParseInt(vals[0], 10, 64)
		n2, _ := strconv.ParseInt(vals[1], 10, 64)
		return []int64{n1, n2}
	})

	mults := lo.Map(pairs, func(inner []int64, _ int) int64 {
		return lo.Reduce(inner, func(acc, item int64, _ int) int64 { return acc * item }, 1)
	})

	sum := lo.Reduce(mults, func(acc, item int64, _ int) int64 { return acc + item }, 0)

	return sum
}

func part2(_ []string) int {
	return 0
}
