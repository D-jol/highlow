package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "djole"
	surname := "title"

	fmt.Printf("Before : %v, %v \n", name, surname)
	//

	name = strings.ToTitle(name)
	surname = strings.ToTitle(surname)
	fmt.Printf("Title : %v, %v \n", name, surname)
	//
	name = strings.ToUpper(name)
	surname = strings.ToUpper(surname)
	fmt.Printf("Uppercase : %v, %v \n", name, surname)
}
