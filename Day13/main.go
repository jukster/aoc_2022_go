package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/jukster/aoc_2022_go/aocutil"
)

func main() {

	input := aocutil.ReadFile("input.txt")

	packetsPartOne := [][]*packet{}

	for i := 0; i < len(input); i += 3 {
		packetOne := parse(input[i])
		packetTwo := parse(input[i+1])
		packetsPartOne = append(packetsPartOne, []*packet{packetOne, packetTwo})
	}
	indices := 0

	for i, p := range packetsPartOne {

		res := compare(p[0], p[1])

		if res == Correct {
			indices += i + 1
		}
	}

	fmt.Println(indices)

	fmt.Println("Part two")

	inputPartTwo := aocutil.ReadFile("input.txt")

	packetsPartTwo := []*packet{}

	for _, line := range inputPartTwo {
		if line != "" {
			packet := parse(line)
			packetsPartTwo = append(packetsPartTwo, packet)
		}
	}

	d1 := packet{
		divider: true,
		packets: []packet{
			{packets: []packet{{value: 2}},
				value: -1}},
		value: -1}

	packetsPartTwo = append(packetsPartTwo, &d1)

	d2 := packet{divider: true, packets: []packet{{packets: []packet{{value: 6}}, value: -1}}, value: -1}

	packetsPartTwo = append(packetsPartTwo, &d2)

	sort.Slice(packetsPartTwo, func(i, j int) bool {
		res := compare(packetsPartTwo[i], packetsPartTwo[j])
		if res == Correct {
			return true
		} else {
			return false
		}
	})

	indicesPartTwo := []int{}

	for i, p := range packetsPartTwo {

		if p.divider {
			indicesPartTwo = append(indicesPartTwo, i+1)
		}

	}

	fmt.Println(indicesPartTwo[0] * indicesPartTwo[1])

}

func compare(left, right *packet) comparisonResult {

	if left.value != -1 {
		left.packets = []packet{{value: left.value}}
	}
	if right.value != -1 {
		right.packets = []packet{{value: right.value}}
	}

	// both are ints - this is only hit as part of the loop below, so
	// inconclusive is allowed, because it is not the end result.

	if left.value != -1 && right.value != -1 {
		return compareInts(left.value, right.value)
	}

	// at least one is list - we need to compare lists
	// we need to iterate as many times as both lists have members

	var i int

	for i < len(left.packets) && i < len(right.packets) {
		if c := compare(&left.packets[i], &right.packets[i]); c != Inconclusive {
			return c
		}
		i++
	}

	// we compared the same ones, still inconclusive, then length decides

	return compareInts(len(left.packets), len(right.packets))

}

type comparisonResult int

const (
	Correct comparisonResult = iota
	Wrong
	Inconclusive
)

func (s comparisonResult) String() string {
	switch s {
	case Correct:
		return "Correct"
	case Wrong:
		return "Wrong"
	case Inconclusive:
		return "Inconclusive"
	}
	return "unknown"
}

func compareInts(left int, right int) comparisonResult {
	var res comparisonResult
	if left == right {
		res = Inconclusive
	} else {
		if left < right {
			res = Correct
		} else {
			res = Wrong
		}
	}
	return res

}

type packet struct {
	divider bool
	packets []packet
	value   int
}

func (p packet) String() string {
	if p.value == -1 {
		return fmt.Sprintf("Packet is a container for %s", p.packets)
	} else {
		return fmt.Sprintf("Packet has value %d", p.value)
	}

}

func (pd *packet) UnmarshalJSON(data []byte) error {
	// custom implementation of unmarshalJSON
	if len(data) == 0 {
		return nil
	}

	if data[0] == '[' {
		// processing the json list
		// making a packet
		// the packet structure contains an attribute with a list of itself as well as the value, so suited for the source data
		var packets []packet
		if err := json.Unmarshal(data, &packets); err != nil {
			return err
		}
		*pd = packet{packets: packets, value: -1}
		return nil
	}

	// function is being called on a single element, we emit the single integer

	var i int
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	*pd = packet{value: i}

	return nil
}

func parse(line string) *packet {
	var v *packet
	if err := json.Unmarshal([]byte(line), &v); err != nil {
		panic(err)
	}
	return v
}
