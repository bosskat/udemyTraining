package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		log.Fatalln(err)
		//		panic(err)
	}
}

/*
package log implements a simple logging package ... writes to standard error
and prints the date and time of each logged message ...
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1)
