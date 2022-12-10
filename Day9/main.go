package main

import (
	"fmt"
	"strconv"

	"github.com/jukster/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	h := newSnake(1)

	tailMoves := []position{}

	for _, line := range inputRaw {
		direction, steps := parseLine(line)
		for i := 0; i < steps; i++ {
			tailMoves = append(tailMoves, h.moveHead(direction))
		}

	}

	fmt.Println(len(unique(tailMoves)))

	fmt.Println("Starting Part 2")

	h2 := newSnake(9)

	tailMoves2 := []position{}

	for _, line := range inputRaw {
		direction, steps := parseLine(line)
		for i := 0; i < steps; i++ {
			tailMoves2 = append(tailMoves2, h2.moveHead(direction))
		}

	}

	fmt.Println(len(unique(tailMoves2)))

}

func parseLine(line string) (direction string, steps int) {
	direction = string(line[0])

	if stepsConverted, err := strconv.Atoi(line[2:]); err == nil {
		steps = stepsConverted

	}

	return direction, steps
}

func unique(s []position) []position {
	inResult := make(map[position]bool)
	var result []position
	for _, pos := range s {
		if _, ok := inResult[pos]; !ok {
			inResult[pos] = true
			result = append(result, pos)
		}
	}
	return result
}

type position struct {
	x int
	y int
}

type snake struct {
	positions []position
}

func (t *snake) moveHead(direction string) (tailPosition position) {
	head := &t.positions[0]
	switch direction {
	case "U":
		head.y--
	case "D":
		head.y++
	case "R":
		head.x++
	case "L":
		head.x--
	default:
		fmt.Println(fmt.Errorf("invalid direction"))
	}

	for i := 1; i < len(t.positions); i++ {
		t.tailChase(&t.positions[i-1], &t.positions[i])
	}

	return t.positions[len(t.positions)-1]
}

func (t *snake) tailChase(leader *position, follower *position) {

	xdelta := leader.x - follower.x
	ydelta := leader.y - follower.y

	moveThreshold := 1

	if xdelta != 0 && ydelta != 0 {
		moveThreshold = 2
	}

	if int(aocutil.Abs(int64(xdelta))+aocutil.Abs(int64(ydelta))) > moveThreshold {

		if xdelta > 0 {
			follower.x++
		}

		if xdelta < 0 {
			follower.x--
		}

		if ydelta > 0 {
			follower.y++
		}

		if ydelta < 0 {
			follower.y--
		}
	}

}

func newSnake(tailLength int) (newSnake snake) {
	newSnake = snake{[]position{{0, 0}}}

	for i := 0; i < tailLength; i++ {
		newSnake.positions = append(newSnake.positions, position{0, 0})
	}

	return newSnake

}
