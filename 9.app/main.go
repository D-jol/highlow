package main

import (
	"fmt"
)

func main() {
	phone_book := make(map[string]string) //made empty array of contact object - phone book

	for {
		fmt.Print("Welcome to phone adress book app\n")
		fmt.Println("[1]- Add contact")
		fmt.Println("[2]- Remove contact")
		fmt.Println("[3]- View contacts")
		fmt.Println("[0]- Exit\n---(Navigate by typing index)---")

		var input int
		fmt.Scanln(&input)
		if input == 0 {
			break
		} else if input == 3 {
			if len(phone_book) == 0 {
				fmt.Println("No contacts")
				break
			} else {
				i := 1
				for key, value := range phone_book {
					fmt.Printf("%v. %v: %v\n", i, key, value)
					i++
				}
			}
		} else if input == 1 {
			// add contact
			var name string
			var phone string
			fmt.Println("Enter name and phone number: ")
			fmt.Scanln(&name, &phone)
			phone_book[name] = phone
		} else if input == 2 {
			i := 1
			for key, value := range phone_book {
				fmt.Printf("%v. %v: %v\n", i, key, value)
				i++
			}
			var del string
			fmt.Println("Type number of the contact you want to delete: ")
			fmt.Scanln(&del)
			for key, value := range phone_book {
				if del == value {
					delete(phone_book, key)
				}
			}
		} else {
			fmt.Println("Invalid input")
			continue
		}

	}
}
