package main

import (
	"reflect"
	"testing"

	"github.com/jukster/aoc_2022_go/aocutil"
)

func TestPartOne(t *testing.T) {

	type scenario struct {
		in   []string
		want int
	}

	inputRawTest := aocutil.ReadFile("input_test2.txt")

	scenarios := []scenario{
		{
			inputRawTest,
			13140,
		},
	}

	for _, sc := range scenarios {
		got := partOne(sc.in)

		if got != sc.want {
			t.Errorf("got %v, wanted %v for %v", got, sc.want, sc.in)
		}
	}
}

func TestPartTwo(t *testing.T) {

	type scenario struct {
		in   []string
		want screen
	}

	inputRawTest := aocutil.ReadFile("input_test2.txt")

	scenarios := []scenario{
		{
			inputRawTest,
			screen{

				[6][40]string{{"#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", "."},
					{"#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", "."},
					{"#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", "."},
					{"#", "#", "#", "#", "#", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", ".", ".", ".", ".", "."},
					{"#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#"},
					{"#", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", "."},
				},
				pixel{0, 0},
			},
		},
	}

	for _, sc := range scenarios {
		got := partTwo(sc.in)

		if !reflect.DeepEqual(got, sc.want) {
			t.Errorf("got \n%vwanted \n%v", got, sc.want)
		}
	}
}
