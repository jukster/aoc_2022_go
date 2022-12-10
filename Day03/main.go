package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jukster/aocutil"
)

func main() {

	inputRaw := aocutil.ReadFile("input.txt")

	prios := 0

	for _, line := range inputRaw {
		comp1, comp2 := splitLine(line)
		common := findCommon(comp1, comp2)

		prios += convertToPriority(common)

	}

	fmt.Println(prios)

	fmt.Println("Starting Day2")

	prios2 := 0

	for i := 0; i < 298; i = i + 3 {

		common := findCommonInThree(
			inputRaw[i],
			inputRaw[i+1],
			inputRaw[i+2],
		)

		prios2 += convertToPriority(common)

	}

	fmt.Println(prios2)

}

func findCommonInThree(first string, second string, third string) string {
	common := ""

	for _, char := range first {
		chars := string(char)
		if strings.Contains(second, chars) {
			if strings.Contains(third, chars) {
				common = chars
			}
		}
	}

	if common == "" {
		panic(errors.New("no common char found"))
	}

	return common

}

func findCommon(first string, second string) string {
	common := ""

	for _, char := range first {
		chars := string(char)
		if strings.Contains(second, chars) {
			common = chars
		}
	}

	if common == "" {
		panic(errors.New("no common char found"))
	}

	return common

}

func splitLine(in string) (partOne string, partTwo string) {
	l := len(in)

	i := l / 2

	return in[:i], in[i:]
}

func convertToPriority(in string) int {
	r := []rune(in)

	ascii := int(r[0])

	if ascii > 96 {
		return ascii - 96
	} else {
		return ascii - 38
	}

}
