package main

import (
	"fmt"
	"math"
)

func solveGrid(grid *grid, startY int, startX int) (steps int) {

	grid.initial = grid.grid[startY][startX]

	dijkstra(grid)

	return grid.pathFind()
}

func main() {

	path := "input.txt"

	fmt.Println("Part 1")

	grid := constructGrid(path)

	fmt.Println(solveGrid(&grid, grid.initial.y, grid.initial.x))

	fmt.Println("Part 2")

	var lowPoints []node

	for _, row := range grid.grid {
		for _, node := range row {
			if node.value == "a" {
				lowPoints = append(lowPoints, *node)
			}
		}
	}

	shortestPath := int(math.Inf(1))

	for _, p := range lowPoints {
		grid := constructGrid(path)
		shortest := solveGrid(&grid, p.y, p.x)
		if shortest < shortestPath {
			shortestPath = shortest
		}
	}

	fmt.Println(shortestPath)

}

func dijkstra(grid *grid) {

	var currentNode *node

	q := queue{}

	currentNode = grid.initial

	for currentNode != grid.final {

		neighbours := currentNode.getAccesibleNeighbours(grid.grid)

		for _, n := range neighbours {
			if currentNode.distanceFromStart+1 < n.distanceFromStart {
				n.distanceFromStart = currentNode.distanceFromStart + 1
				n.parent = currentNode
				q.add(n)
			}
		}

		currentNode.visited = true

		currentNode = q.getNext()

		if currentNode == nil {
			break
		}

	}

}
