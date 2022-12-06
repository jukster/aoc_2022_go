package main

import (
	"testing"
)

func TestCheckDistinct(t *testing.T) {

	type scenario struct {
		input string
		want  bool
	}

	scenarios := []scenario{
		{
			"abcb",
			false,
		},
		{
			"abcd",
			true,
		},
	}

	for _, sc := range scenarios {
		got := checkRepeat(sc.input)

		if got == sc.want {
			t.Errorf("2: got %v, wanted %v", got, sc.want)
		}
	}

}
