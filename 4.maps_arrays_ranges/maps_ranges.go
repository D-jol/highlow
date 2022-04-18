package main

import "fmt"

func main() {

	m := make(map[string]string)

	m["djole"] = "help desk"
	m["a"] = "123"
	m["b"] = "456"

	//fmt.Println(m["djole"])

	for key, item := range m {
		if key == "a" {
			continue
		}
		fmt.Println(key, item)
	}

}
