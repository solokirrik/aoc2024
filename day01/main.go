package main

import (
	_ "embed"
	"log/slog"
	"sort"
	"strconv"
	"strings"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1", "Diff", part1(inp))
	slog.Info("Part 2", "Diff", part2(inp))
}

func part1(inp string) int {
	lists := toSortedLists(inp)
	diff := 0

	for i := range lists[0] {
		val := max(lists[0][i], lists[1][i]) - min(lists[0][i], lists[1][i])
		diff += val
	}

	return diff
}

func part2(inp string) int {
	lists := toSortedLists(inp)
	diff := 0

	for _, item := range lists[0] {
		sim := countSims(item, lists[1])
		diff += item * sim
	}

	return diff
}

func countSims(target int, exp []int) int {
	sim := 0

	for _, val := range exp {
		if val == target {
			sim++
		}
	}

	return sim
}

func toSortedLists(in string) [2][]int {
	lines := strings.Split(in, "\n")

	out := [2][]int{make([]int, 0, len(lines)), make([]int, 0, len(lines))}

	for i := range lines {
		elems := strings.Fields(lines[i])

		lVal, err := strconv.Atoi(elems[0])
		panicOnErr(err)

		rVal, err := strconv.Atoi(elems[1])
		panicOnErr(err)

		out[0] = append(out[0], lVal)
		out[1] = append(out[1], rVal)
	}

	sort.Ints(out[0])
	sort.Ints(out[1])

	return out
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
