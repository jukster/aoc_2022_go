package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jukster/aoc_2022_go/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	indices := extractStackIndices(inputRaw)

	stacks := parseInitialStacks(inputRaw, indices)

	stacks2 := copyDeep(stacks)

	// Processing Part 1

	for _, line := range inputRaw[10:] {

		inst := newInstruction(line)

		removed := stacks[inst.from].removeStacks(inst.quantity)

		for i := len(removed) - 1; i >= 0; i-- {
			stacks[inst.to].crates = prependStr(stacks[inst.to].crates, removed[i])
		}

	}

	solution := ""

	for _, stack := range stacks {
		solution = solution + stack.crates[0]
	}

	fmt.Println(solution)

	fmt.Println("Starting Day2")

	for _, line := range inputRaw[10:] {

		inst := newInstruction(line)

		removed := stacks2[inst.from].removeStacks(inst.quantity)

		for i := 0; i < len(removed); i++ {
			stacks2[inst.to].crates = prependStr(stacks2[inst.to].crates, removed[i])
		}

	}

	solution2 := ""

	for _, stack := range stacks2 {
		solution2 = solution2 + stack.crates[0]
	}

	fmt.Println(solution2)

}

type instruction struct {
	quantity int
	from     int
	to       int
}

func newInstruction(line string) instruction {

	numsPat := regexp.MustCompile(`\d+`)
	matches := numsPat.FindAllStringSubmatch(line, -1)

	converted := []int{}

	for _, str := range matches {
		t, err := strconv.Atoi(str[0])

		if err != nil {
			fmt.Errorf("conversion failed")
		}

		converted = append(converted, t)

	}

	q, f, t := converted[0], converted[1], converted[2]

	f = f - 1

	t = t - 1

	return instruction{q, f, t}
}

func parseInitialStacks(inputRaw []string, indices []int) (stacks [9]stack) {

	for i, index := range indices {
		stackStr := []string{}
		for x := 0; x < 8; x++ {
			char := string(inputRaw[x][index])
			if char != " " {
				stackStr = append(stackStr, char)
			}
		}
		stacks[i] = stack{stackStr}
	}

	return stacks
}

func copyDeep(orig [9]stack) (copied [9]stack) {

	for i, s := range orig {
		copied[i].crates = make([]string, len(s.crates))
		copy(copied[i].crates, s.crates)
	}

	return copied
}

type stack struct {
	crates []string
}

func (s *stack) removeStacks(howMany int) (removed []string) {

	for i := 0; i < howMany; i++ {

		removed = prependStr(removed, s.crates[i])

	}

	for i := 0; i < howMany; i++ {
		if len(s.crates) > 0 {
			s.crates = s.crates[1:]
		}
	}

	return removed

}

func extractStackIndices(input []string) (indices []int) {

	theLine := input[8]

	for i, c := range theLine {
		if string(c) != " " {
			indices = append(indices, i)
		}
	}
	return indices
}

func prependStr(x []string, y string) []string {
	x = append(x, "")
	copy(x[1:], x)
	x[0] = y

	return x
}
