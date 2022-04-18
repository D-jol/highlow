package main

import (
	"fmt"
)

var (
	user_info string
)

func user_info_input() (user_info string) {
	var name, surname string

	fmt.Print("What's your name: ")
	fmt.Scanln(&name)
	fmt.Print("What's your surname: ")
	fmt.Scanln(&surname)

	user_info = name + " " + surname

	return
}

func main() {
	user := user_info_input()
	fmt.Println("Hello, ", user, "!")

}
