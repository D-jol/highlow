package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name and age here(don't use commas) :")

	fmt.Scan(&name, &age)
	fmt.Printf("Name: %v | Age: %v", name, age)
}
