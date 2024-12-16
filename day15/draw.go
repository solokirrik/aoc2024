package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"time"
)

var printSleep = 1 * time.Millisecond

const (
	bold = "\033[1m"

	green           = "\033[32m"
	reset           = "\033[0m"
	black           = "\033[30m"
	red             = "\033[31m"
	yellow          = "\033[33m"
	blue            = "\033[34m"
	magenta         = "\033[35m"
	cyan            = "\033[36m"
	white           = "\033[37m"
	brightBlackGray = "\033[90m"
	brightRed       = "\033[91m"
	brightGreen     = "\033[92m"
	brightYellow    = "\033[93m"
	brightBlue      = "\033[94m"
	brightMagenta   = "\033[95m"
	brightCyan      = "\033[96m"
	brightWhite     = "\033[97m"
)

func (s *solver) fprint(i int) {
	fmt.Println(i)
	fmt.Println(s.stringify())
}

func (s *solver) stringify() string {
	var out string

	for _, row := range s.mtx {
		idx := slices.Index(row, "@")
		if idx != -1 {
			out += brightBlackGray + strings.Join(row[:idx], "") + reset +
				yellow + bold + "@" + reset +
				brightBlackGray + strings.Join(row[idx+1:], "") + reset + "\n"
		} else {
			out += brightBlackGray + strings.Join(row, "") + reset + "\n"
		}
	}

	return out
}

func (s *solver) print(i int) {
	if !s.doPrint {
		return
	}

	if s.doSleep {
		time.Sleep(printSleep)
	}

	if s.doClean {
		clean(len(s.mtx))
	}

	fmt.Println(i)
	fmt.Println(s.stringify())
}

func (s *solver) fieldIsBroken() bool {
	var out string
	for _, row := range s.mtx {
		out += strings.Join(row, "") + "\n"
	}

	return strings.Contains(out, "[.") ||
		strings.Contains(out, ".]") ||
		strings.Contains(out, "[#") ||
		strings.Contains(out, "#]") ||
		strings.Contains(out, "[[") ||
		strings.Contains(out, "]]")
}

const ESC = 27

var clear = fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)

func clean(linesY int) {
	_, _ = fmt.Fprint(io.Writer(os.Stdout), strings.Repeat(clear, linesY+2))
}
