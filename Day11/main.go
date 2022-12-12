package main

import (
	"fmt"
	"math"
	"sort"
)

func solve(monkeys []monkey, partTwo bool) (productOfTwo int) {

	rounds := 20

	if partTwo {
		rounds = 10000
	}

	for _, monkey := range monkeys {
		monkey.counter = len(monkey.items)
	}
	var msp []*monkey

	for i := range monkeys {
		msp = append(msp, &monkeys[i])
	}

	for i := 0; i < rounds; i++ {

		playRound(msp, partTwo)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].counter > monkeys[j].counter
	})

	return monkeys[0].counter * monkeys[1].counter
}

func main() {

	fmt.Println("part 1")

	fmt.Println(solve(readMonkeys("input.txt"), false))

	fmt.Println("part 2")

	fmt.Println(solve(readMonkeys("input.txt"), true))

}

type monkey struct {
	items           []item
	operation       Operation
	testDivisibleBy int
	throwTrue       int
	throwFalse      int
	counter         int
}

func (m monkey) String() string {

	retStr := fmt.Sprintf("Monkey (counter %d)with items: %v\n", m.counter, m.items)

	return retStr
}

type item struct {
	worry int
}

func (i *item) reduceWorry() {
	i.worry = int(math.Floor(float64(i.worry) / 3))
}

func (i *item) evaluateWorry(divideBy int) bool {
	if i.worry%divideBy == 0 {
		return true
	} else {
		return false
	}
}

type Operation func(int) int

func playRound(monkeys []*monkey, partTwo bool) {

	lcm := 1

	for _, monkey := range monkeys {
		lcm = lcm * monkey.testDivisibleBy
	}

	for i := 0; i < len(monkeys); i++ {
		for x := 0; x < len(monkeys[i].items); x++ {
			it := monkeys[i].items[x]
			it.worry = monkeys[i].operation(it.worry)

			if partTwo {
				it.worry = it.worry % lcm
			} else {
				it.reduceWorry()
			}
			var monkeyReceives int
			if it.evaluateWorry(monkeys[i].testDivisibleBy) {
				monkeyReceives = monkeys[i].throwTrue
			} else {
				monkeyReceives = monkeys[i].throwFalse
			}

			monkeys[monkeyReceives].items = append(monkeys[monkeyReceives].items, it)
			monkeys[i].counter++

		}

		monkeys[i].items = []item{}
	}
}
