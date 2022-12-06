package main

import "fmt"

func main() {

	arr := []astruct{
		astruct{2},
		astruct{3},
		astruct{4},
	}

	arrc := arr

	fmt.Println(arr)
	fmt.Println(arrc)

	arr[0].value = 5

	fmt.Println(arr)
	fmt.Println(arrc)

}

type astruct struct {
	value int
}
