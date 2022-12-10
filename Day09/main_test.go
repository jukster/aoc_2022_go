package main

import (
	"testing"
)

func TestMoveHead(t *testing.T) {

	type scenario struct {
		in        snake
		want      position
		direction string
	}

	scenarios := []scenario{
		{
			snake{
				[]position{
					{2, 5},
					{1, 5},
				},
			},
			position{2, 5},
			"R",
		},
		{
			snake{
				[]position{
					{3, 5},
					{2, 5},
				},
			},
			position{2, 5},
			"U",
		},
		{
			snake{
				[]position{
					{3, 4},
					{2, 5},
				},
			},
			position{3, 4},
			"U",
		},
	}

	for _, sc := range scenarios {
		got := sc.in.moveHead(sc.direction)

		if got != sc.want {
			t.Errorf("got %v, wanted %v for %v", got, sc.want, sc.in)
		}
	}
}
