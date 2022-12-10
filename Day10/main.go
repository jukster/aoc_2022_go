package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jukster/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	fmt.Println(partOne(inputRaw))

	fmt.Println("Starting Part 2")

	fmt.Println(partTwo(inputRaw))

}

type screen struct {
	grid         [6][40]string
	currentPixel pixel
}

func (s screen) String() string {

	retStr := ""

	for _, row := range s.grid {
		retStr += strings.Join(row[:], "")
		retStr += "\n"
	}

	return retStr
}

func (s *screen) drawPixel(lit bool) {
	if lit {
		s.grid[s.currentPixel.y][s.currentPixel.x] = "#"
	} else {
		s.grid[s.currentPixel.y][s.currentPixel.x] = "."
	}
	if s.currentPixel.x == 39 {
		s.currentPixel.x = 0
		if s.currentPixel.y == 5 {
			s.currentPixel.y = 0
		} else {
			s.currentPixel.y += 1
		}
	} else {
		s.currentPixel.x += 1
	}

}

type pixel struct {
	x int
	y int
}

func isSpriteOverlapping(x int, pixel int) bool {
	if pixel <= x+1 && pixel >= x-1 {
		return true
	} else {
		return false
	}
}

func partTwo(inputRaw []string) screen {

	s := screen{}

	cycle := 1

	xValues := []int{1}

	for _, line := range inputRaw {
		noop, addx := parseLine(line)

		pixelPosition := (cycle % 40) - 1

		if noop {
			s.drawPixel(isSpriteOverlapping(xValues[len(xValues)-1], pixelPosition))
			cycle += 1
		} else {
			s.drawPixel(isSpriteOverlapping(xValues[len(xValues)-1], pixelPosition))
			cycle += 1
			pixelPosition++
			s.drawPixel(isSpriteOverlapping(xValues[len(xValues)-1], pixelPosition))
			xValues = append(xValues, addx+xValues[len(xValues)-1])
			cycle += 1
		}

	}

	return s
}

func partOne(inputRaw []string) (total int) {

	significantSignals := map[int]int{}

	significantCycles := []int{20, 60, 100, 140, 180, 220}

	cycle := 1

	xValues := []int{1}

	for _, line := range inputRaw {
		noop, addx := parseLine(line)

		for _, s := range significantCycles {
			if cycle == s {
				if _, ok := significantSignals[s]; !ok {
					significantSignals[s] = xValues[len(xValues)-1] * s
				}
			} else if cycle > s {
				if _, ok := significantSignals[s]; !ok {
					significantSignals[s] = xValues[len(xValues)-2] * s
				}
			}
		}

		if noop {
			cycle += 1
		} else {
			cycle += 2
			xValues = append(xValues, addx+xValues[len(xValues)-1])
		}

	}

	for _, value := range significantSignals {
		total += value
	}

	return total
}

func parseLine(line string) (noop bool, addx int) {
	if strings.Contains(line, "noop") {
		noop = true
	} else {
		if num, err := strconv.Atoi(string(line[5:])); err == nil {
			addx = num
		}
	}

	return noop, addx
}
