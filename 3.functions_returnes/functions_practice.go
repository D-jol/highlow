package main

import (
	"fmt"
)

var (
	user_info string
)

func user_info_input() string {
	var name, surname string

	fmt.Print("name: ")
	fmt.Scanln(&name)
	fmt.Print("surname: ")
	fmt.Scanln(&surname)

	user_info = name + " " + surname

	return user_info
}

func main() {
	user := user_info_input()
	fmt.Println(user)

	for range user {
		fmt.Println(user)
	}
}
