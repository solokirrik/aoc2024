package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"strings"
)

//go:embed ex
var inp string

func main() {
	inplines := strings.Split(inp, "\n")

	slog.Info("Part 1:", "safe:", part1(inplines))
	slog.Info("Part 2:", "safe:", part2(inplines))
}

// 123 4 567 8 91011 12
// mul ( ... ,  ...  )
const mul = "mul("

func part1(inp []string) int {
	result := 0
	for _, row := range inp {
		result += calcMuls(row)
	}

	return result
}

func calcMuls(row string) int {
	result := 0

	tmp := row

	for strings.Contains(tmp, mul) {
		pos := strings.Index(tmp, mul)
		rPos := strings.Index(tmp[pos:pos+12], ")")

		if rPos == -1 {
			tmp = tmp[pos+rPos+1:]
			continue
		}

		comma := strings.Index(tmp[pos+4:pos+rPos], ",")
		if comma == -1 {
			tmp = tmp[pos+rPos+1:]
			continue
		}

		// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
		fmt.Println(tmp[pos+4:pos+4+comma], tmp[pos+4+comma+1:pos+4+rPos])

		tmp = tmp[rPos+1:]
	}

	return result
}

func part2(_ []string) int {
	return 0
}
