package main

import (
	"strconv"
	"strings"

	"github.com/jukster/aocutil"
)

func readMonkeys(path string) (monkeys []monkey) {
	input := aocutil.ReadFile(path)

	for i := 0; i < len(input); i += 7 {
		var items []item
		for _, nums := range strings.Split(input[i+1][18:], ", ") {
			num, _ := strconv.Atoi(nums)
			items = append(items, item{num})
		}

		operation := input[i+2][23:]

		operationFun := genFunc(string(operation[0]), string(operation[2:]))

		divideBy, _ := strconv.Atoi(input[i+3][21:])

		throwTrue, _ := strconv.Atoi(input[i+4][29:])

		throwFalse, _ := strconv.Atoi(input[i+5][30:])

		monkeys = append(monkeys, monkey{items, operationFun, divideBy, throwTrue, throwFalse, 0})

	}

	return monkeys
}

func genFunc(operator string, value string) (fun func(int) int) {
	num, err := strconv.Atoi(value)

	if err != nil {
		fun = func(old int) int { return old * old }
	} else {

		switch operator {
		case "*":
			fun = func(old int) int { return old * num }
		case "+":
			fun = func(old int) int { return old + num }
		}
	}

	return fun

}
