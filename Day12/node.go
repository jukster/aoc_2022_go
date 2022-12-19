package main

import (
	"fmt"
	"math"
)

type node struct {
	value             string
	y                 int
	x                 int
	visited           bool
	distanceFromStart int
	parent            *node
}

func (n node) String() string {
	return fmt.Sprintf("Node [%d, %d], value %v ", n.y, n.x, n.value)
}

func newNode(value string, y int, x int) node {

	node := node{value, y, x, false, int(math.Inf(1)), nil}

	return node
}

func (n *node) getAccesibleNeighbours(grid [][]*node) (an []*node) {
	if n.y > 0 {
		dest_up := grid[n.y-1][n.x]
		if canCross(n, dest_up) && !dest_up.visited {
			an = append(an, dest_up)
		}
	}

	if n.y < len(grid)-1 {
		dest_down := grid[n.y+1][n.x]
		if canCross(n, dest_down) && !dest_down.visited {
			an = append(an, dest_down)
		}
	}

	if n.x > 0 {
		dest_left := grid[n.y][n.x-1]
		if canCross(n, dest_left) && !dest_left.visited {
			an = append(an, dest_left)
		}
	}

	if n.x < len(grid[0])-1 {
		dest_right := grid[n.y][n.x+1]
		if canCross(n, dest_right) && !dest_right.visited {
			an = append(an, dest_right)
		}
	}

	return an
}

func canCross(from *node, to *node) bool {
	fromA := int(from.value[0])
	toA := int(to.value[0])
	if fromA-toA > -2 {
		return true
	} else {
		return false
	}
}
