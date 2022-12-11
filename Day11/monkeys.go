package main

func makeMonkeys() []monkey {
	monkeys := []monkey{
		{
			[]item{
				{
					71,
				},
				{
					56,
				},
				{
					50,
				}, {
					73,
				}},
			func(old int) int {
				return old * 11
			},
			13,
			1,
			7,
			0,
		},
		{
			[]item{
				{
					70,
				},
				{
					89,
				},
				{
					82,
				},
			},
			func(old int) int {
				return old + 1
			},
			7,
			3,
			6,
			0,
		},
		{
			[]item{
				{
					52,
				},
				{
					95,
				},
			},
			func(old int) int {
				return old * old
			},
			3,
			5,
			4,
			0,
		},
		{
			[]item{
				{
					94,
				},
				{
					64,
				}, {
					69,
				}, {
					87,
				}, {
					70,
				}},
			func(old int) int {
				return old + 2
			},
			19,
			2,
			6,
			0,
		},
		{
			[]item{
				{
					98,
				},
				{
					72,
				}, {
					98,
				}, {
					53,
				}, {
					97,
				}, {
					51,
				}},
			func(old int) int {
				return old + 6
			},
			5,
			0,
			5,
			0,
		},
		{
			[]item{
				{
					79,
				}},
			func(old int) int {
				return old + 7
			},
			2,
			7,
			0,
			0,
		},
		{
			[]item{
				{
					77,
				},
				{
					55,
				}, {
					63,
				}, {
					93,
				}, {
					66,
				}, {
					90,
				}, {
					88,
				}, {
					71,
				}},
			func(old int) int {
				return old * 7
			},
			11,
			2,
			4,
			0,
		},
		{
			[]item{
				{
					54,
				},
				{
					97,
				}, {
					87,
				}, {
					70,
				}, {
					59,
				}, {
					82,
				}, {
					59,
				}},
			func(old int) int {
				return old + 8
			},
			17,
			1,
			3,
			0,
		},
	}

	return monkeys
}
