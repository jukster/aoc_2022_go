package main

import (
	"fmt"
	"strings"

	"github.com/jukster/aoc_2022_go/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	fmt.Println(processStream(inputRaw[0], 4))

	fmt.Println("Day 2")

	fmt.Println(processStream(inputRaw[0], 14))

}

func checkRepeat(input string) bool {
	repeat := false

	for _, char := range input {
		c := string(char)

		if strings.Count(input, c) > 1 {
			repeat = true
		}
	}

	return repeat
}

func processStream(input string, strLength int) (markerIndex int) {

	for i := 0; i < len(input)-strLength; i++ {

		str := input[i : i+strLength]

		check := checkRepeat(str)

		if !check {
			markerIndex = i + strLength
			break
		}

	}
	return markerIndex
}
