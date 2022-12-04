package main

import (
	"testing"
)

func TestCheckOverLap(t *testing.T) {

	type scenario struct {
		in   []aRange
		want overlapResult
	}

	scenarios := []scenario{
		{
			[]aRange{
				{2, 4},
				{6, 8},
			},
			overlapResult{false, false, false},
		},
		{
			[]aRange{
				{2, 3},
				{4, 5},
			},
			overlapResult{false, false, false},
		},
		{
			[]aRange{
				{5, 7},
				{7, 9},
			},
			overlapResult{true, false, false},
		},
		{
			[]aRange{
				{2, 8},
				{3, 7},
			},
			overlapResult{true, false, true},
		},
		{
			[]aRange{
				{6, 6},
				{4, 6},
			},
			overlapResult{true, true, false},
		},
		{
			[]aRange{
				{2, 6},
				{4, 8},
			},
			overlapResult{true, false, false},
		},
	}

	for _, sc := range scenarios {
		r := sc.in[0]
		got := r.checkOverlap(sc.in[1])

		if got != sc.want {
			t.Errorf("got %v, wanted %v", got, sc.want)
		}
	}

}
