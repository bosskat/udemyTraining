package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"
	greeting = append(greeting, "Suprabadham")

	fmt.Println(greeting[3])
}
