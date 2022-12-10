package main

import (
	"fmt"
	"strconv"

	"github.com/jukster/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	grid := makeGrid(99, 99)

	for y, line := range inputRaw {
		for x, char := range line {
			int, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(fmt.Errorf("conversion failed"))
			}
			grid[y][x] = int
		}
	}

	fmt.Println(countVisible(grid))

	fmt.Println("Part 2")

	fmt.Println(countScenicScore(grid))

}

func makeGrid(width int, height int) (grid [][]int) {

	for i := 0; i < height; i++ {
		row := make([]int, width)
		grid = append(grid, row)
	}

	return grid
}

func countScenicScore(grid [][]int) (highestScenicScore int) {
	for y, row := range grid {
		for x := range row {
			directions := [][]int{
				row[x:],
				reverse(row[:x+1]),
				getVerticalUp(grid, x, y),
				getVerticalDown(grid, x, y),
			}

			thisScenicScore := 1

			for _, d := range directions {

				thisScenicScore = thisScenicScore * viewDistance(d)
			}

			if thisScenicScore > highestScenicScore {
				highestScenicScore = thisScenicScore
			}

		}
	}
	return highestScenicScore

}

func countVisible(grid [][]int) (countOfVisible int) {
	for y, row := range grid {
		for x := range row {
			directions := [][]int{
				row[x:],
				reverse(row[:x+1]),
				getVerticalUp(grid, x, y),
				getVerticalDown(grid, x, y),
			}

			treeVisible := false

			for _, d := range directions {
				if isFirstMax(d) {
					treeVisible = true
				}
			}

			if treeVisible {
				countOfVisible++
			}

		}
	}
	return countOfVisible

}

func viewDistance(arr []int) (distance int) {

	breakLoop := false

	distance = 0

	for _, tree := range arr[1:] {

		if tree >= arr[0] {
			breakLoop = true
		}
		distance++
		if breakLoop {
			break
		}

	}

	return distance

}

func isFirstMax(arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	isFirstMax := true

	for _, num := range arr[1:] {
		if num >= arr[0] {
			isFirstMax = false
		}
	}

	return isFirstMax
}

func reverse(in []int) []int {
	r := []int{}

	for i := len(in) - 1; i >= 0; i-- {

		r = append(r, in[i])

	}

	return r
}

func getVerticalUp(grid [][]int, x int, y int) (arr []int) {
	for y := y; y >= 0; y-- {
		arr = append(arr, grid[y][x])
	}

	return arr
}

func getVerticalDown(grid [][]int, x int, y int) (arr []int) {
	for y := y; y < len(grid); y++ {
		arr = append(arr, grid[y][x])
	}

	return arr
}
