package main

import (
	"fmt"
)

func main() {
	var array [3]string // array with 3 values

	fmt.Print("Enter some names :")
	fmt.Scan(&array[0], &array[1], &array[2])

	for _, value := range array {
		fmt.Println(value)
	}

}
