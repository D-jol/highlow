package main

import "fmt"

func main() {

	var array [3]string

	array[0] = "pivo"
	array[1] = "sok"
	array[2] = "cips"
	for _, value := range array {
		fmt.Println(value)
	}

}
