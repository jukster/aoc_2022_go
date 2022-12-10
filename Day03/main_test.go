package main

import (
	"testing"
)

func TestSplitLine(t *testing.T) {
	in := "vJrwpWtwJgWrhcsFMMfFFhFp"

	got1, got2 := splitLine(in)

	want1 := "vJrwpWtwJgWr"

	want2 := "hcsFMMfFFhFp"

	if got1 != want1 {
		t.Errorf("got %v, wanted %v", got1, want1)
	}

	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}

}

func TestConvertToPriority(t *testing.T) {
	type scenario struct {
		in   string
		want int
	}
	scenarios := []scenario{
		{
			"p",
			16,
		},
		{
			"L",
			38,
		},
		{
			"P",
			42,
		},
		{
			"v",
			22,
		},
		{
			"t",
			20,
		},
		{
			"s",
			19,
		},
	}

	for _, sc := range scenarios {
		got := convertToPriority(sc.in)

		if got != sc.want {
			t.Errorf("got %v, wanted %v", got, sc.want)
		}
	}

}

func TestFindCommonInThree(t *testing.T) {
	in := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}

	got := findCommonInThree(in[0], in[1], in[2])

	want := "r"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
