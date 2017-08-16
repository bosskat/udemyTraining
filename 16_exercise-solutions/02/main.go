package main

import "fmt"

func main() {
	fmt.Println(half(6))
	half := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}
	fmt.Println(half(18))
	fmt.Println(greatest(3, 7, 2, 9))
	fmt.Println(boolean())
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func half(x int) (int,bool) {
	return x / 2, x%2 == 0
}

func greatest(x ...int) int {
	var biggest int
	for _, i := range x {
		if i > biggest {
			biggest = i
		}
	}
	return biggest
}

func boolean() bool {
	return (true && false) || (false && true) || !(false && false)
}

func foo(x ...int) {
	fmt.Println(x)
}
