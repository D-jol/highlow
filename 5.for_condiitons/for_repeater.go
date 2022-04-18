package main

import (
	"fmt"
)

func main() {

	var repeat int
	var word string

	fmt.Println("Tell me a word: ")
	fmt.Scan(&word)
	fmt.Println("How many times to repeat ?: ")
	fmt.Scan(&repeat)

	for ; repeat != 0; repeat-- {
		fmt.Println("\n", word)
	}

}
