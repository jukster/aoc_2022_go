package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jukster/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	calorieRegistry := make(map[int]int)

	elfIndex := 0

	for _, line := range inputRaw {

		if line == "" {
			elfIndex++
		} else {
			num, err := strconv.Atoi(line)

			if err != nil {
				fmt.Println(fmt.Errorf("conversion failed"))

			}

			calorieRegistry[elfIndex] = calorieRegistry[elfIndex] + num
		}

	}

	winner := 0

	for _, calorieCount := range calorieRegistry {
		if calorieCount > winner {
			winner = calorieCount
		}
	}

	fmt.Println(winner)

	fmt.Println("Starting Day 2")

	counts := make([]int, len(calorieRegistry))

	for i, count := range calorieRegistry {
		counts[i] = count
	}

	sort.Ints(counts)

	topThree := 0

	for i := len(counts) - 3; i < len(counts); i++ {
		topThree = topThree + counts[i]
	}

	fmt.Println(topThree)
}
