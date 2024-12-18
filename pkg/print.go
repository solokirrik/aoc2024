package pkg

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const ESC = 27

var clear = fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)

func Clear(linesY int) {
	_, _ = fmt.Fprint(io.Writer(os.Stdout), strings.Repeat(clear, linesY+2))
}

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
