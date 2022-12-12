package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {

	got := solve(readMonkeys("input_test.txt"), false)

	want := 10605

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got \n%vwanted \n%v", got, want)
	}

}
