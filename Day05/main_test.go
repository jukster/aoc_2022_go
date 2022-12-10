package main

import (
	"reflect"
	"testing"
)

func TestRemoveStacks(t *testing.T) {

	type scenario struct {
		stack     stack
		howMany   int
		want      []string
		wantStack stack
	}

	scenarios := []scenario{
		{stack{
			[]string{"N", "Z"}},
			1,
			[]string{"N"},
			stack{[]string{"Z"}},
		},
		{stack{
			[]string{"D", "N", "Z"},
		},
			3,
			[]string{"Z", "N", "D"},
			stack{[]string{}},
		},
		{stack{
			[]string{"Z", "G", "J", "P", "Q", "D", "L", "M"},
		},
			8,
			[]string{"M", "L", "D", "Q", "P", "J", "G", "Z"},
			stack{[]string{}},
		},
	}

	for _, sc := range scenarios {
		got := sc.stack.removeStacks(sc.howMany)

		if !reflect.DeepEqual(got, sc.want) {
			t.Errorf("got %v, wanted %v", got, sc.want)
		}

		if !reflect.DeepEqual(sc.stack, sc.wantStack) {
			t.Errorf("2: got %v, wanted %v", sc.stack, sc.wantStack)
		}
	}

}
