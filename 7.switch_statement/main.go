package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	access := false

	fmt.Print("Name validation: \n")
	fmt.Scan(&name)
	name = strings.ToLower(name)

	switch name {
	case "djole", "john", "doe":
		fmt.Print("welcome")
		access = true
	default:
		fmt.Println("Non authorised user.")
	}

	if access {
		fmt.Print("\nYou can enter the chat")
	} else {
		fmt.Print("\nCalling the cops !")
	}

}
