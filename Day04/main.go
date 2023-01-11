package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jukster/aoc_2022_go/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	countContainments := 0

	countOverlaps := 0

	for _, line := range inputRaw {

		firstSplit := strings.Split(line, ",")

		ranges := []aRange{}

		for _, elf := range firstSplit {

			secondSplit := strings.Split(elf, "-")

			boundaries := []int{}

			for _, num := range secondSplit {

				converted, err := strconv.Atoi(num)

				if err != nil {
					fmt.Errorf("conversion failed")
				}

				boundaries = append(boundaries, converted)

			}

			ranges = append(ranges, aRange{boundaries[0], boundaries[1]})
		}

		res := ranges[0].checkOverlap(ranges[1])

		if res.initialContained || res.otherContained {
			countContainments++
		}

		if res.overlaps {
			countOverlaps++
		}
	}

	fmt.Println(countContainments)

	fmt.Println("Starting Day2")

	fmt.Println(countOverlaps)

}

type aRange struct {
	start int
	end   int
}

type overlapResult struct {
	overlaps         bool
	initialContained bool
	otherContained   bool
}

func (r *aRange) checkOverlap(otherRange aRange) (result overlapResult) {

	overlaps := false

	initialContained := false

	otherContained := false

	if r.start >= otherRange.start && r.end <= otherRange.end {
		initialContained = true
	}

	if r.start <= otherRange.start && r.end >= otherRange.end {
		otherContained = true
	}

	if r.start >= otherRange.start && r.start <= otherRange.end || r.start <= otherRange.start && r.end >= otherRange.start {
		overlaps = true
	}

	return overlapResult{overlaps, initialContained, otherContained}
}
