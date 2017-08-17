package main

import "fmt"

func main() {
	//helloWorld()
	//myNameIs()
	//yourNameIs()
	//remainder()
	//even()
	//fizzBuzz()
	multiples()
}

func helloWorld() {
	fmt.Println("Hello World")
}

func myNameIs() {
	name := "bosskat"
	fmt.Println("Hello, my name is", name)
}

func yourNameIs() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}

func remainder() {
	var x, y int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&x)
	fmt.Print("Enter a larger number: ")
	fmt.Scan(&y)
	fmt.Println(y % x)
}

func even() {
	for n := 0; n <= 100; n++ {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}
}

func fizzBuzz() {
	for n := 1; n <= 100; n++ {
		if n%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if n%3 == 0 {
			fmt.Println("Fizz")
		} else if n%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(n)
		}
	}
}

func multiples() {
	tot := 0
	for i := 0; i < 1000; i++ {

		if i%3 == 0 {
			tot += i
		} else if i%5 == 0 {
			tot += i
		}
	}
	fmt.Println(tot)
}
