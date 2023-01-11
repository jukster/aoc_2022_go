package main

import (
	"math"
	"strconv"

	"github.com/jukster/aoc_2022_go/aocutil"
)

type grid struct {
	grid    [][]*node
	initial *node
	final   *node
}

func (g *grid) pathFind() (steps int) {

	var currentNode *node

	currentNode = g.final

	for currentNode != g.initial {
		if currentNode.parent == nil {
			return int(math.Inf(1))
		}
		currentNode = currentNode.parent
		steps++
	}

	return steps
}

func constructGrid(path string) (newGrid grid) {

	newGrid = grid{
		[][]*node{},
		nil,
		nil,
	}

	input := aocutil.ReadFile(path)

	height := len(input)
	width := len(input[0])

	for i := 0; i < height; i++ {
		row := make([]*node, width)
		newGrid.grid = append(newGrid.grid, row)
	}

	for y, line := range input {
		for x, str := range line {
			s := string(str)
			thisNode := newNode(s, y, x)
			if s == "S" {
				newGrid.initial = &thisNode
				thisNode.distanceFromStart = 0
				thisNode.value = "a"
			}

			if s == "E" {
				newGrid.final = &thisNode
				thisNode.value = "z"
			}

			newGrid.grid[y][x] = &thisNode
		}
	}

	return newGrid
}

func (g grid) String() string {
	retStr := ""

	for _, row := range g.grid {
		rowstr := ""
		for _, node := range row {
			if node.visited {
				rowstr = rowstr + "|" + strconv.Itoa(node.distanceFromStart) + "|"
			} else {
				rowstr = rowstr + "."
			}
		}

		retStr = retStr + rowstr + "\n"
	}

	return retStr
}
