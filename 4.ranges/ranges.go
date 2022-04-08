package main

import "fmt"

func main() {
	/*
		v := "hello"
		/*
		for _, c := range v {
			fmt.Println(string(c))
		}
	*/
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
