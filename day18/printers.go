package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func print(mtx [][]uint8, visited map[string]int, newPos coord) {
	var out string

	for r := range mtx {
		newRow := make([]string, len(mtx[r]))
		for c := range mtx[r] {
			newRow[c] = black + "." + reset

			if mtx[r][c] == 1 {
				newRow[c] = brightBlackGray + "#" + reset
				continue
			}

			curPos := coord{r: r, c: c}
			if ok := visitedContains(curPos, visited); ok {
				newRow[c] = green + "0" + reset + brightBlackGray
			}

			if r == newPos.r && c == newPos.c {
				newRow[c] = bold + red + "X" + reset
			}
		}

		out += strings.Join(newRow, "") + "\n"
	}

	clean(len(mtx))
	fmt.Println(len(visited), newPos.r, newPos.c)
	fmt.Println(out)
}

func printWithBroken(mtx [][]uint8, visited map[string]int, broken []coord) {
	var out string

	for r := range mtx {
		newRow := make([]string, len(mtx[r]))
		for c := range mtx[r] {
			newRow[c] = black + "." + reset

			if mtx[r][c] == 1 {
				newRow[c] = brightBlackGray + "#" + reset
				continue
			}

			curPos := coord{r: r, c: c}
			if ok := visitedContains(curPos, visited); ok {
				newRow[c] = green + "0" + reset + brightBlackGray
			}

			if isInCoords(curPos, broken) {
				newRow[c] = bold + red + "X" + reset
			}
		}

		out += strings.Join(newRow, "") + "\n"
	}

	clean(len(mtx))
	fmt.Println(out)
}

func isInCoords(c coord, coll []coord) bool {
	for _, s := range coll {
		if c.eq(s) {
			return true
		}
	}

	return false
}

func dVisitedContains(c coord, vis map[string]int) bool {
	for _, dir := range []string{NORTH, EAST, SOUTH, WEST} {
		dc := dcoord{c: c, dir: dir}
		if _, ok := vis[dc.hash()]; ok {
			return true
		}
	}

	return false
}

func visitedContains(c coord, vis map[string]int) bool {
	if _, ok := vis[c.hash()]; ok {
		return true
	}

	return false
}

const (
	bold  = "\033[1m"
	reset = "\033[0m"

	green           = "\033[32m"
	black           = "\033[30m"
	red             = "\033[31m"
	brightBlackGray = "\033[90m"
	brightGreen     = "\033[92m"
)

const ESC = 27

var clear = fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)

func clean(linesY int) {
	_, _ = fmt.Fprint(io.Writer(os.Stdout), strings.Repeat(clear, linesY+2))
}
