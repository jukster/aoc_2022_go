package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	partTwo := false

	ms := makeMonkeys()

	for _, monkey := range ms {
		monkey.counter = len(monkey.items)
	}
	var msp []*monkey

	for i := range ms {
		msp = append(msp, &ms[i])
	}

	rounds := 20

	if partTwo {
		rounds = 10000
	}

	for i := 0; i < rounds; i++ {

		playRound(msp, partTwo)
	}

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].counter > ms[j].counter
	})

	fmt.Println(ms[0].counter * ms[1].counter)

	fmt.Println("Starting part 2")

	partTwo = true

	ms2 := makeMonkeys()

	for _, monkey := range ms2 {
		monkey.counter = len(monkey.items)
	}
	var msp2 []*monkey

	for i := range ms2 {
		msp2 = append(msp2, &ms2[i])
	}

	if partTwo {
		rounds = 10000
	}

	for i := 0; i < rounds; i++ {

		playRound(msp2, partTwo)
	}

	sort.Slice(ms2, func(i, j int) bool {
		return ms2[i].counter > ms2[j].counter
	})

	fmt.Println(ms2[0].counter * ms2[1].counter)

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
