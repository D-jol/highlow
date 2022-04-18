package main

import (
	"fmt"
)

func main() {
	num1, num2 := 5, 100

	if num1 == 5 || num2 == 10 {
		fmt.Print("correct\n")
	} else {
		fmt.Println("false")
	}
	if num1 == 5 && num2 == 10 {
		fmt.Print("correct")
	} else {
		fmt.Println("false")
	}

}
