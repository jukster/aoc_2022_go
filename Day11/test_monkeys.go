package main

func makeTestMonkeys() []monkey {
	monkeys := []monkey{
		{
			[]item{
				{
					79,
				},
				{
					98,
				},
			},
			func(old int) int {
				return old * 19
			},
			23,
			2,
			3,
			0,
		},
		{
			[]item{
				{
					54,
				},
				{
					65,
				},
				{
					75,
				},
				{
					74,
				},
			},
			func(old int) int {
				return old + 6
			},
			19,
			2,
			0,
			0,
		},
		{
			[]item{
				{
					79,
				},
				{
					60,
				},
				{
					97,
				},
			},
			func(old int) int {
				return old * old
			},
			13,
			1,
			3,
			0,
		},
		{
			[]item{
				{
					74,
				},
			},
			func(old int) int {
				return old + 3
			},
			17,
			0,
			1,
			0,
		},
	}

	return monkeys
}
