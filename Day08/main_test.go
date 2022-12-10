package main

import "testing"

func TestIsFirstMax(t *testing.T) {
	type scenario struct {
		in   []int
		want bool
	}

	scenarios := []scenario{
		{
			[]int{3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 1},
			false,
		},
		{
			[]int{0},
			true,
		},
		{
			[]int{2, 0, 0},
			true,
		},
	}

	for _, sc := range scenarios {
		got := isFirstMax(sc.in)

		if got != sc.want {
			t.Errorf("got %v, wanted %v for %v", got, sc.want, sc.in)
		}
	}
}

func TestCountVisible(t *testing.T) {
	type scenario struct {
		grid [][]int
		want int
	}

	scenarios := []scenario{
		{
			[][]int{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 9, 9},
			},
			21,
		},
	}

	for _, sc := range scenarios {
		got := countVisible(sc.grid)

		if got != sc.want {
			t.Errorf("got %v, wanted %v for %v", got, sc.want, sc.grid)
		}
	}

}

func TestCountScenicScore(t *testing.T) {
	type scenario struct {
		grid [][]int
		want int
	}

	scenarios := []scenario{
		{
			[][]int{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 9, 9},
			},
			8,
		},
	}

	for _, sc := range scenarios {
		got := countScenicScore(sc.grid)

		if got != sc.want {
			t.Errorf("got %v, wanted %v for %v", got, sc.want, sc.grid)
		}
	}

}

func TestViewDistance(t *testing.T) {
	type scenario struct {
		in   []int
		want int
	}

	scenarios := []scenario{
		{
			[]int{3, 2, 1},
			2,
		},
		{
			[]int{1, 2, 1},
			1,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{2, 0, 0, 1, 2},
			4,
		},
		{
			[]int{2, 5, 3, 5},
			1,
		},
		{
			[]int{3, 0, 3, 7, 3},
			2,
		},
	}

	for _, sc := range scenarios {
		got := viewDistance(sc.in)

		if got != sc.want {
			t.Errorf("got %v, wanted %v for %v", got, sc.want, sc.in)
		}
	}
}
