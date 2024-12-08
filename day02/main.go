package main

import (
	_ "embed"
	"errors"
	"log/slog"
	"strconv"
	"strings"

	"github.com/solokirrik/aoc2024/utils"
)

//go:embed inp
var inp string

var errUnsafe = errors.New("line is unfase")

const (
	incr = "incr"
	decr = "decr"
)

func main() {
	inplines := strings.Split(inp, "\n")

	slog.Info("Part 1:", "safe:", part1(inplines))
	slog.Info("Part 2:", "safe:", part2(inplines))
}

func part1(inpLines []string) int {
	safe := 0

	for _, line := range inpLines {
		report := getReport(line)

		if err := isSafe(report); err != nil {
			continue
		}

		safe++
	}

	return safe
}

func part2(inpLines []string) int {
	safe := 0

	for _, line := range inpLines {
		wasSafe := false
		report := getReport(line)

		for _, mut := range mutations(report) {
			if err := isSafe(mut); err == nil {
				wasSafe = true
				break
			}
		}

		if wasSafe {
			safe++
		}
	}

	return safe
}

func getReport(inp string) []int {
	chars := strings.Split(inp, " ")
	report := make([]int, 0, len(chars))

	for i := range chars {
		c1, _ := strconv.Atoi(chars[i])
		report = append(report, c1)
	}

	return report
}

func mutations(report []int) [][]int {
	mutations := [][]int{report}

	for i := range len(report) {
		option := make([]int, 0, len(report)-1)

		for j := range len(report) {
			if i != j {
				option = append(option, report[j])
			}
		}

		mutations = append(mutations, option)
	}

	return mutations
}

func isSafe(report []int) error {
	oldDir, err := toDir(report[0], report[1])
	if err != nil {
		return err
	}

	for i := 0; i < len(report)-1; i++ {
		if err := check(report[i], report[i+1], oldDir); err != nil {
			return err
		}
	}

	return nil
}

func toDir(a, b int) (string, error) {
	switch {
	case a == b:
		return "", errUnsafe
	case a > b:
		return decr, nil
	case a < b:
		return incr, nil
	default:
		return "", nil
	}
}

func check(a, b int, oldDir string) error {
	if utils.Abs(a-b) < 1 || utils.Abs(a-b) > 3 {
		return errUnsafe
	}

	dir, err := toDir(a, b)
	if err != nil {
		return errUnsafe
	}

	if oldDir != dir {
		return errUnsafe
	}

	return nil
}
