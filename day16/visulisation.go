package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/solokirrik/aoc2024/utils"
)

var (
	blueP  = color.RGBA{173, 216, 230, 255}
	greenP = color.RGBA{0, 255, 0, 255}
	gray   = color.RGBA{128, 128, 128, 255}
)

func (s *solver) writeGif(images []*image.Paletted, delays []int) {
	f, err := os.Create("output-" + strconv.FormatInt(time.Now().Unix(), 10) + ".gif")
	utils.PanicOnErr(err)

	defer f.Close()

	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func (s *solver) savePaths(bestScore int, winners []step) {
	seenPaths := make(map[string]bool, len(winners))
	for i := range winners {
		if winners[i].score == bestScore {
			pHash := pathHash(winners[i].path)
			if seenPaths[pHash] {
				continue
			}
			seenPaths[pHash] = true
			s.render(i, winners[i].path)
		}
	}
}

func (s *solver) renderVisited(i int64) {
	width, height := len(s.mtx[0]), len(s.mtx)
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	file, err := os.Create("./pics/" + strconv.FormatInt(i, 10) + "path.png")
	utils.PanicOnErr(err)

	defer file.Close()

	for r, row := range s.mtx {
		for c := range row {
			curPos := coord{r: r, c: c}
			img.Set(c, r, gray)
			if s.mtx[r][c] == "#" {
				img.Set(c, r, color.Black)
			}
			if s.visited[curPos.hash()] > 0 {
				img.Set(c, r, greenP)
			}
		}
	}

	utils.PanicOnErr(png.Encode(file, img))
}

func (s *solver) render(i int, path []coord) {
	img := image.NewRGBA(image.Rect(0, 0, len(s.mtx[0]), len(s.mtx)))

	file, err := os.Create("./pics/" + strconv.FormatInt(int64(i), 10) + "path.png")
	utils.PanicOnErr(err)

	defer file.Close()

	for r, row := range s.mtx {
		for c := range row {
			curPos := coord{r: r, c: c}
			img.Set(c, r, gray)
			if s.mtx[r][c] == "#" {
				img.Set(c, r, color.Black)
			}
			if idx := slices.Index(path, curPos); idx != -1 {
				img.Set(c, r, greenP)
			}
		}
	}

	utils.PanicOnErr(png.Encode(file, img))
}

func (s *solver) visitedFrame() *image.Paletted {
	palette := []color.Color{color.Black, blueP, gray, greenP}

	width, height := len(s.mtx[0]), len(s.mtx)
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	for r, row := range s.mtx {
		for c := range row {
			img.Set(c, r, gray)
			if s.mtx[r][c] == "#" {
				img.Set(c, r, color.Black)
			}
		}
	}

	for k := range s.visited {
		pos := deDhash(k)
		img.Set(pos.c, pos.r, greenP)
	}

	return img
}

func deDhash(in string) coord {
	parts := strings.Split(in, "-")
	r, _ := strconv.ParseInt(parts[1], 10, 64)
	c, _ := strconv.ParseInt(parts[2], 10, 64)
	return coord{r: int(r), c: int(c)}
}

func (s *solver) printSeats(seats map[string]coord) {
	var out string

	for r, row := range s.mtx {
		for c := range row {
			curPos := coord{r: r, c: c}
			if _, ok := seats[curPos.hash()]; ok {
				row[c] = green + "O" + reset + brightBlackGray
			}
		}
		newRow := brightBlackGray + strings.Join(row, "") + reset
		out += newRow + "\n"
	}

	fmt.Println(out)
}

func dpathContains(dpath []dcoord, p coord) (string, bool) {
	for i := range dpath {
		if dpath[i].getPos().eq(p) {
			return dpath[i].dir, true
		}
	}
	return "", false
}

func (s *solver) print(dpath []dcoord) {
	var out string

	for r, row := range s.mtx {
		for c := range row {
			curPos := coord{r: r, c: c}
			if dir, ok := dpathContains(dpath, curPos); ok {
				row[c] = green + dir + reset + brightBlackGray
			}
		}
		newRow := brightBlackGray + strings.Join(row, "") + reset
		out += newRow + "\n"
	}

	fmt.Println(out)
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
